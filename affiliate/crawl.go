package affiliate

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/4meepo/tiktok-tools/elegant"
	"github.com/4meepo/tiktok-tools/ent"
	"github.com/4meepo/tiktok-tools/ent/tiktokcreator"
)

func CrawlAffiliateCreators(curlSample, region string, duration time.Duration) error {
	ctx, cancelFn := context.WithCancel(context.Background())
	ec := ent.GetInstance()

	go elegant.Shutdown(cancelFn)

	var followerCntMax = getMinFollowerCount(region)

	// 此api每个请求最多返回2000条数据, 我们每爬1000条就休息一段时间,下次重新开始爬取
	const size = 100
	const maxPagePerBatch = 10 // 每次最多爬取10页

	var retryTimes, totalUpdate, totalInsert int
	for page := 0; ; page++ {
		// 休息一段时间
		if page == maxPagePerBatch {
			log.Printf("防止封号 休息 %s 后继续... \n", duration.String())
			time.Sleep(duration)
			page = 0                                     // 重新开始
			followerCntMax = getMinFollowerCount(region) // 重新获取最小粉丝数
		}

		var nextItemCursor *int
		if page != 0 {
			nic := (page * size) + 1
			nextItemCursor = &nic
		}
		request := searchCreatorsRequest{
			Request: requestPayload{
				Algorithm:      3,
				FollowerCntMax: followerCntMax,
				Pagination: pagination{
					Size:           size,
					Page:           page,
					NextItemCursor: nextItemCursor,
				},
			},
		}

		rsp, err := searchCreators(curlSample, request)
		if err != nil {
			if errors.Is(err, ErrCurlCmd) {
				return err
			}
			log.Printf("搜索达人失败: %s\n", err.Error())
			retry(&page, &retryTimes)
			continue
		}
		if rsp.Code != 0 {
			log.Printf("搜索达人失败:%d %s\n", rsp.Code, rsp.Message)
			retry(&page, &retryTimes)
			continue
		}

		if rsp.Data.NextPagination.HasMore == false {
			log.Printf("无更多达人信息,搜索达人结束\n")
			os.Exit(1)
		}

		// 先查询是否重复
		var creatorsId []string
		for _, c := range rsp.Data.CreatorProfiles {
			creatorsId = append(creatorsId, c.CreatorId)
		}
		creators, err := ec.TiktokCreator.Query().Where(tiktokcreator.CreatorIDIn(creatorsId...)).All(ctx)
		if err != nil {
			log.Printf("查询达人失败: %s\n", err.Error())
			retry(&page, &retryTimes)
			continue
		}
		creatorsMap := make(map[string]struct{})
		for _, c := range creators {
			creatorsMap[c.CreatorID] = struct{}{}
		}
		for _, c := range rsp.Data.CreatorProfiles {
			select {
			case <-ctx.Done():
				os.Exit(0)
			default:
				// 重复的更新
				if _, ok := creatorsMap[c.CreatorId]; ok {
					_, err := ec.TiktokCreator.Update().Where(tiktokcreator.CreatorIDEQ(c.CreatorId)).
						SetCreatorName(c.CreatorName).
						SetCreatorNickname(c.CreatorNickName).
						SetRegion(region).
						SetProductCategories(c.ProductCategories).
						SetFollowerCount(c.FollowerCount).
						SetVideoAvgViewCnt(c.VideoAvgViewCnt).
						SetVideoPubCnt(c.VideoPubCnt).
						SetEcVideoAvgViewCnt(c.EcVideoAvgViewCnt).
						SetCreatorOecuid(c.CreatorOecuid).
						SetCreatorTtuid(c.CreatorTtuid).Save(ctx)
					if err != nil {
						log.Printf("更新达人失败: %s\n", err.Error())
						continue
					}
					totalUpdate++
					log.Printf("更新达人信息成功: %s %s\n", c.CreatorId, c.CreatorName)
				} else {
					// 不重复的插入
					_, err := ec.TiktokCreator.Create().
						SetCreatorID(c.CreatorId).
						SetCreatorName(c.CreatorName).
						SetCreatorNickname(c.CreatorNickName).
						SetRegion(region).
						SetProductCategories(c.ProductCategories).
						SetFollowerCount(c.FollowerCount).
						SetVideoAvgViewCnt(c.VideoAvgViewCnt).
						SetVideoPubCnt(c.VideoPubCnt).
						SetEcVideoAvgViewCnt(c.EcVideoAvgViewCnt).
						SetCreatorOecuid(c.CreatorOecuid).
						SetCreatorTtuid(c.CreatorTtuid).Save(ctx)
					if err != nil {
						log.Printf("插入达人数据失败: %s\n", err.Error())
						continue
					}
					totalInsert++
					log.Printf("插入达人数据成功: creatodId: %s, userName:%s 粉丝数: %d\n", c.CreatorId, c.CreatorName, c.FollowerCount)
				}
			}
		}

		// 清零
		retryTimes = 0

		// 输出统计信息
		log.Printf("第%d页, 更新%d条, 插入%d条\n", page, totalUpdate, totalInsert)
		time.Sleep(randomDuration())

	}
}

func retry(page, retryTimes *int) {
	if *retryTimes >= 5 {
		log.Printf("重试次数超过5次, 停止爬取\n")
		os.Exit(1)
	}
	*page--
	*retryTimes++
}

func randomDuration() time.Duration {
	min, max := 15, 45
	s := min + rand.Intn(max-min)
	return time.Duration(s) * time.Second
}

func getMinFollowerCount(region string) *uint32 {
	ec := ent.GetInstance()
	// 从数据库中获取最小的粉丝数, 作为下一次的最大粉丝数
	lastC, err := ec.TiktokCreator.Query().
		Where(tiktokcreator.RegionEQ(region)).
		Order(ent.Asc(tiktokcreator.FieldFollowerCount)).
		First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		} else {
			log.Printf("查询最后一个达人失败: %s", err)
			os.Exit(1)
			return nil
		}
	} else {
		return &lastC.FollowerCount
	}
}
