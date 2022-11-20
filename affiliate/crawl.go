package affiliate

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/4meepo/tiktok-tools/elegant"
	"github.com/4meepo/tiktok-tools/ent"
	"github.com/4meepo/tiktok-tools/ent/tiktokcreator"
)

var httpClient = http.Client{
	Timeout: time.Second * 8,
}

func CrawlAffiliateCreators(host,
	curlSample string,
	duration time.Duration, // 休息时长
	affiliateFollowerFrom, // 从多少粉丝开始爬取, 间隔为100
	pageSize,
	threshold int, // 最多爬取多少粉丝后休息一段时间
) error {
	ctx, cancelFn := context.WithCancel(context.Background())
	ec := ent.GetInstance(host)

	go elegant.Shutdown(cancelFn)

	var followerCntMax = affiliateFollowerFrom

	var retryTimes, totalUpdate, totalInsert int
	for page, requestCount := 0, 1; ; page++ {
		select {
		case <-ctx.Done():
			os.Exit(1)
		default:
			//
		}
		// 休息一段时间
		if requestCount%threshold == 0 {
			log.Printf("防止封号 休息 %s 后继续... \n", duration.String())
			time.Sleep(duration)
		}

		var nextItemCursor *int
		if page != 0 {
			nic := (page * pageSize) + 1
			nextItemCursor = &nic
		}
		request := searchCreatorsRequest{
			Request: requestPayload{
				Algorithm:      3,
				FollowerCntMax: followerCntMax,
				FollowerCntMin: followerCntMax - 100,
				Pagination: pagination{
					Size:           pageSize,
					Page:           page,
					NextItemCursor: nextItemCursor,
				},
			},
		}

		rsp, err := searchCreators(curlSample, request)
		requestCount++
		if err != nil {
			if errors.Is(err, ErrCurlCmd) {
				return err
			}
			log.Printf("搜索达人失败: %v\n", err)
			retry(&page, &retryTimes)
			continue
		}
		if rsp.Code != 0 {
			log.Printf("搜索达人失败:%d msg:%s\n", rsp.Code, rsp.Message)
			retry(&page, &retryTimes)
			continue
		}

		if rsp.Data.NextPagination.HasMore == false {
			// 此区间无更多达人
			// next round
			log.Printf("[%d %d]区间无更多达人, 下一个区间\n", followerCntMax, followerCntMax-100)
			followerCntMax -= 100
			if followerCntMax <= 100 {
				log.Printf("粉丝数已经小于100, 退出爬取")
				os.Exit(0)
			}
			page = -1 // 重新开始
		} else {

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
							SetRegion(c.Region).
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
							SetRegion(c.Region).
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
			log.Printf("第%d页, 更新%d条, 插入%d条, 区间[%d %d]共有 %d 人 \n", page, totalUpdate, totalInsert, followerCntMax, followerCntMax-100, rsp.Data.NextPagination.Total)
		}

		d := randomDuration()
		log.Printf("休息 %s 后继续\n", d.String())
		time.Sleep(d)
	}
}

func retry(page, retryTimes *int) {
	if *retryTimes >= 10 {
		log.Printf("连续失败次数超过10次, 停止爬取\n")
		os.Exit(1)
	}
	*page--
	*retryTimes++
}

func randomDuration() time.Duration {
	min, max := 15, 30
	s := min + rand.Intn(max-min)
	return time.Duration(s) * time.Second
}
