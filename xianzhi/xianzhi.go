package xianzhi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/4meepo/tiktok-tools/elegant"
	"github.com/4meepo/tiktok-tools/ent"
	"github.com/4meepo/tiktok-tools/ent/creator"
)

func CrawlCreators(host, xianzhiRegion, authorization, userId string, fromPage, batchSize int, duration time.Duration) error {
	rand.Seed(time.Now().Unix())

	ctx, cancelFn := context.WithCancel(context.Background())
	go elegant.Shutdown(cancelFn)

	entClient := ent.GetInstance(host)

	var size = 100
	var _count = 0
	for i := fromPage; ; i++ {
		select {
		case <-ctx.Done():
			log.Println("程序退出...")
			return nil
		default:
			log.Printf("已爬取 %d 条数据,正在爬取第 %d 页数据...", _count, i)
			rsp, err := queryByRegion(xianzhiRegion, userId, authorization, i, size)
			if err != nil {
				log.Printf("获取第 %d 页数据失败: [%v]\n", i, err)
				i--
				continue
			}
			if len(rsp.Data.List) == 0 {
				// 结束
				log.Printf("获取第 %d 页数据为空, 爬取结束,准备退出 \n", i)
				os.Exit(0)
			}

			// 剔除重复数据
			xzids := make([]string, 0, batchSize)
			for _, v := range rsp.Data.List {
				xzids = append(xzids, v.Xzid)
			}
			creators, err := entClient.Creator.Query().Where(creator.XzidIn(xzids...)).All(ctx)
			if err != nil {
				log.Printf("查询重复数据失败: [%v]\n", err)
				i--
				continue
			}
			// to map
			creatorMap := make(map[string]struct{}, len(creators))
			for _, v := range creators {
				creatorMap[v.Xzid] = struct{}{}
			}

			// batch insert to db
			bulk := make([]*ent.CreatorCreate, 0, len(rsp.Data.List))
			for _, item := range rsp.Data.List {
				if _, ok := creatorMap[item.Xzid]; ok {
					log.Printf("数据重复,跳过: %s\n", item.Xzid)
					continue
				}
				c := entClient.Creator.Create().
					SetXzid(item.Xzid).
					SetUniqueID(item.UniqueID).
					SetNickName(item.NickName).
					SetRegion(item.Region).
					SetFollowerNum(uint32(item.FollowerNum)).
					SetCate1NameCn(item.ProductCategories).
					SetEmail(item.Email).
					SetWhatsapp(item.Whatsapp)
				bulk = append(bulk, c)
				_count++
			}
			if err := entClient.Creator.CreateBulk(bulk...).Exec(ctx); err != nil {
				log.Printf("第 %d 页数据插入数据库失败, 准备重新爬取, reason: [%v]\n", i, err)
				i--
				continue
			}

		}
		// sleep
		time.Sleep(randomDuration())

		// 休息
		if i%batchSize == 0 {
			log.Printf("防止封号 休息 %s \n", duration.String())
			time.Sleep(duration)
		}
	}
}

func queryByRegion(region, userId, authorization string, page, pageSize int) (*XianzhiBatchResponse, error) {
	host := "https://usermgr.xianzhiai.com/search/kol/categorySearch"

	httpRequest, err := http.NewRequest("POST", host, strings.NewReader(fmt.Sprintf(`
{
    "page": %d,
    "pageSize": %d,
    "userId": "%s",
    "canContact": null,
    "categoryId": null,
    "followerCountEnd": null,
    "followerCountStart": 10000,
    "region": "%s",
    "sortType": "followerNum",
	"sortWay": "desc",
    "keyWord": ""
}
`, page, pageSize, userId, region)))

	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorization))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: [%w]", err)
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: [%w]", err)
	}
	defer rsp.Body.Close()

	bytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: [%w]", err)
	}
	var batchResponse XianzhiBatchResponse
	if err := json.Unmarshal(bytes, &batchResponse); err != nil {
		return nil, fmt.Errorf("解析响应失败: [%w]", err)
	}
	return &batchResponse, nil
}

type XianzhiBatchResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type Item struct {
	Xzid              string   `json:"xzid"`
	UniqueID          string   `json:"unique_id"`
	NickName          string   `json:"nick_name"`
	Region            string   `json:"region"`
	FollowerNum       int      `json:"follower_num"`
	ProductCategories []string `json:"cate1_name_cn"`
	Email             string   `json:"email"`
	Whatsapp          string   `json:"whatsapp"`
}
type Data struct {
	Count int    `json:"count"`
	List  []Item `json:"dataList"`
}

func randomDuration() time.Duration {
	min, max := 30, 60
	s := min + rand.Intn(max-min)
	return time.Duration(s) * time.Second
}
