/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"log"
)

// xianzhiCmd represents the xianzhi command
var xianzhiCmd = &cobra.Command{
	Use:   "xianzhi",
	Short: "从 先知网 爬取达人数据",
	RunE: func(cmd *cobra.Command, args []string) error {
		return crawlXianzhiCreators()
	},
}

var authorization, xianzhiOutputFile, xianzhiRegion, xianzhiCategory string

var cookie = `_ttp=2G3pjDXmtE6k16A1PCS7sEURBSu; passport_csrf_token=71ea17e27a8eea603f728400e3a8d632; passport_csrf_token_default=71ea17e27a8eea603f728400e3a8d632; _ga=GA1.1.1869615319.1667906360; _fbp=fb.1.1667906360498.560041025; _tt_enable_cookie=1; sso_uid_tt_ads=9c5430aa50d4d0647a8a0b0516a5b8bfa04ce51b1d4491758d99b62ad35e4419; sso_uid_tt_ss_ads=9c5430aa50d4d0647a8a0b0516a5b8bfa04ce51b1d4491758d99b62ad35e4419; toutiao_sso_user_ads=4260065941ef762207316225353bf649; toutiao_sso_user_ss_ads=4260065941ef762207316225353bf649; sid_ucp_sso_v1_ads=1.0.0-KGRmMTcxMTFmNWJiNmE0ZjFkN2EyMWRjNmI0ZGZlNjBmZTNlODU4N2YKIAiBiILu_LOXlGMQ0vaomwYY5B8gDDDJ3KKZBjgBQOsHEAEaA3NnMSIgNDI2MDA2NTk0MWVmNzYyMjA3MzE2MjI1MzUzYmY2NDk; ssid_ucp_sso_v1_ads=1.0.0-KGRmMTcxMTFmNWJiNmE0ZjFkN2EyMWRjNmI0ZGZlNjBmZTNlODU4N2YKIAiBiILu_LOXlGMQ0vaomwYY5B8gDDDJ3KKZBjgBQOsHEAEaA3NnMSIgNDI2MDA2NTk0MWVmNzYyMjA3MzE2MjI1MzUzYmY2NDk; ttwid=1%7CJamJ6DSQTUBINetCMp_9f7ZBH8J3PPob_yVRTm6_010%7C1667906389%7C631f0381daa050c040f45aa294dd2a758dfe5035df9542b3796fadae3b655884; SELLER_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcxNDUwNjM3NDk3NTY0MjExMjEsIk9lY1VpZCI6NzQ5NDc1OTAwMTk4NzEyMzYxNSwiT2VjU2hvcElkIjo3NDk0NzU5MDAxOTg3MTIzNjE1LCJTaG9wUmVnaW9uIjoiIiwiR2xvYmFsU2VsbGVySWQiOjc0OTQ3NTkwMDE5ODcxMjM2MTUsIlNlbGxlcklkIjo3NDk0NzU5MDAxOTg3MTIzNjE1LCJleHAiOjE2Njc5OTI3OTMsIm5iZiI6MTY2NzkwNTM5M30.uDPqJ_6rjDfFPojwlMGLZ9s2uqIb8sSr6M11OQzZe1o; SHOP_ID=7151810777085640965; _ga_BZBQ2QHQSP=GS1.1.1667906360.1.1.1667906400.0.0.0; s_v_web_id=verify_la84enq3_EnaXuk5G_tFEe_4XrE_B7YJ_vLn7wn1fz7NJ; uid_tt_ads=9c5430aa50d4d0647a8a0b0516a5b8bfa04ce51b1d4491758d99b62ad35e4419; uid_tt_ss_ads=9c5430aa50d4d0647a8a0b0516a5b8bfa04ce51b1d4491758d99b62ad35e4419; sid_tt_ads=4260065941ef762207316225353bf649; sessionid_ads=4260065941ef762207316225353bf649; sessionid_ss_ads=4260065941ef762207316225353bf649; sid_guard_ads=4260065941ef762207316225353bf649%7C1667906402%7C863985%7CFri%2C+18-Nov-2022+11%3A19%3A47+GMT; sid_guard_tiktokshop=74d6df3fc78d8e1e726dbb52f5a7f37d%7C1667906402%7C863985%7CFri%2C+18-Nov-2022+11%3A19%3A47+GMT; uid_tt_tiktokshop=29008d51666292b2049bc9d2386a5578f81a41ad6f96235f4d49e0f6e91969cc; uid_tt_ss_tiktokshop=29008d51666292b2049bc9d2386a5578f81a41ad6f96235f4d49e0f6e91969cc; sid_tt_tiktokshop=74d6df3fc78d8e1e726dbb52f5a7f37d; sessionid_tiktokshop=74d6df3fc78d8e1e726dbb52f5a7f37d; sessionid_ss_tiktokshop=74d6df3fc78d8e1e726dbb52f5a7f37d; sid_ucp_v1_tiktokshop=1.0.0-KDhiNjdjODEwN2ExZmYwMTkzOGVjODk5YTZlMjA5ZTk1YmU3OGQ2MTIKGgiBiILu_LOXlGMQ4vaomwYY5B8gDDgBQOsHEAMaBm1hbGl2YSIgNzRkNmRmM2ZjNzhkOGUxZTcyNmRiYjUyZjVhN2YzN2Q; ssid_ucp_v1_tiktokshop=1.0.0-KDhiNjdjODEwN2ExZmYwMTkzOGVjODk5YTZlMjA5ZTk1YmU3OGQ2MTIKGgiBiILu_LOXlGMQ4vaomwYY5B8gDDgBQOsHEAMaBm1hbGl2YSIgNzRkNmRmM2ZjNzhkOGUxZTcyNmRiYjUyZjVhN2YzN2Q; i18next=en; csrf_session_id=f2f24c3966162c59f9b0e9ba9aef4e8f; user_oec_info=0a53dea892b2dfb3160664fdf2065dd9fdb98de6bbe2a30092edbe9f5122c13a22f169f2f679f3cf0a98b6a54706d70bd5eed4aeebe114d46ca391c56648c0c42d6e4bc1060499639150f037b0a8ee5c1ba22478131a490a3c4891e7960dd966fa44782efc12ebca1a2aa6a98ad4976846d701137acf138f9478e0212c6643c37286091a6c6ea5bcf4150b6151f4f4b7a677453b861087d6a00d1886d2f6f20d220104e059c2c1; odin_tt=6b464a755499628e17c2c54a0ab2b6bc055cada631dfc9e0b6a05405256a46fe7226c744c55617680a6a0101a91d48c4bad716e3912983aa204d20e1ccf767b6; msToken=mEbJNPywA5fnmzrU4T6MyP3m0uF7Q6lufux6mBxX7c2z_33rYL7juFA4DB4eLvl8WJX9Wn35MWps_PFUTfiBYIkSmTovsDCllPXKP-nwirl_YjDIoyd1CdT6Glklv8fD0fqGUA==; msToken=8EXXMcHd1H0khDY5K9YrrpeysihGT9ulMR3BBJud-1Jr_LqujHRbhCZNZv4CKdDLZfL31w6x_qBE2HXKl5etz7ESkCAapBpBhgGkikcVLL6q_owtqvu6nggjqY8PyQO4usheHtcYaHmiNCU=`

func init() {
	crawlCmd.AddCommand(xianzhiCmd)

	xianzhiCmd.Flags().StringVarP(&authorization, "authorization", "a", "", "先知网的 authorization")
	xianzhiCmd.Flags().StringVarP(&xianzhiOutputFile, "outputFile", "o", "", "输出文件名")
	xianzhiCmd.Flags().StringVarP(&xianzhiRegion, "region", "r", "", "地区 ID--印尼 MY--马来西亚 TH--泰国  GB--英国")
	xianzhiCmd.Flags().StringVarP(&xianzhiCategory, "category", "c", "", "分类 1001--服装 1002--鞋子 1003-配饰 1004--箱包 1005--玩具 1006--美妆个护 1007--运动户外 1008--消费类电子 1009--家居 1010--母婴 1011--汽车/摩托 1012--日用百货 1013--宠物用品")
}

func crawlXianzhiCreators() error {
	outputF, err := os.Create(fmt.Sprintf("%s", strings.TrimSpace(xianzhiOutputFile)))
	if err != nil {
		return fmt.Errorf("创建文件 %s 失败: [%w] ", xianzhiOutputFile, err)
	}
	defer outputF.Close()
	w := csv.NewWriter(outputF)
	w.Write([]string{"creator_id", "creator_name", "creator_nickname", "region", "product_categories", "follower_cnt", "video_avg_view_cnt", "creator_oecuid", "主页"})

	ctx, cancelFn := context.WithCancel(context.Background())
	go elegantShutdown(cancelFn)

	var size = 50
	for i := 1; ; i++ {
		rsp, err := fetchXianzhiNext(xianzhiRegion, xianzhiCategory, i, size)
		if err != nil {
			log.Printf("获取第 %d 页数据失败: [%v]\n", i, err)
			i--
			continue
		}
		if len(rsp.Data.DataList) == 0 {
			// 结束
			log.Printf("获取第 %d 页数据为空, 爬取结束,准备退出 \n", i)
			break
		}
		for _, item := range rsp.Data.DataList {
			select {
			case <-ctx.Done():
				log.Println("程序退出...")
				return nil
			default:
				writeTikTokUserInfo(w, item.UniqueID)
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}

	return nil
}

func fetchXianzhiNext(region, category string, page, pageSize int) (*XianzhiBatchResponse, error) {
	host := "https://usermgr.xianzhiai.com/search/kol/categorySearch"

	httpRequest, err := http.NewRequest("POST", host, strings.NewReader(fmt.Sprintf(`
{
    "page": %d,
    "pageSize": %d,
    "userId": "8a6694c681a5e34f0181a96f8863001c",
    "canContact": null,
    "categoryId": null,
    "followerCountEnd": null,
    "followerCountStart": null,
    "region": "%s",
    "sortType": "recommendRank",
    "keyWord": ""
}
`, page, pageSize, region)))

	httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorization))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: [%w]", err)
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("请求失败: [%w]", err)
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

func writeTikTokUserInfo(w *csv.Writer, uid string) {
	req, err := http.NewRequest("POST", "https://affiliate.tiktok.com/api/v1/oec/affiliate/creator/marketplace/search?user_language=en&aid=4331&app_name=i18n_ecom_alliance&device_id=0&fp=verify_la84enq3_EnaXuk5G_tFEe_4XrE_B7YJ_vLn7wn1fz7NJ&device_platform=web&cookie_enabled=true&screen_width=2048&screen_height=1280&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Mozilla&browser_version=5.0+(Macintosh%3B+Intel+Mac+OS+X+10_15_7)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F107.0.0.0+Safari%2F537.36+Edg%2F107.0.1418.28&browser_online=true&timezone_name=Asia%2FShanghai&shop_region=TH&msToken=8EXXMcHd1H0khDY5K9YrrpeysihGT9ulMR3BBJud-1Jr_LqujHRbhCZNZv4CKdDLZfL31w6x_qBE2HXKl5etz7ESkCAapBpBhgGkikcVLL6q_owtqvu6nggjqY8PyQO4usheHtcYaHmiNCU=&X-Bogus=DFSzswVLJO2PDTaES0aV/5KMtadE&_signature=_02B4Z6wo00001S3uaggAAIDA06IWOY0bs4Et7m6AACgUa1", strings.NewReader(fmt.Sprintf(`
{
    "request": {
        "follower_genders": [],
        "follower_age_groups": [],
        "managed_by_agency": [],
        "pagination": {
            "size": 20,
            "page": 0
        },
        "creator_score_range": [],
        "content_preference_range": [],
        "algorithm": 1,
        "query": "%s"
    }
}

	`, uid)))
	if err != nil {
		log.Printf("创建tiktok 查询用户信息请求失败: [%v]\n", err)
		return
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/json")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("请求tiktok 查询用户信息失败: [%v]\n", err)
		return
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		log.Printf("请求tiktok 查询用户信息失败: [%v]\n", rsp.Status)
		return
	}

	bytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("读取tiktok 查询用户信息响应失败: [%v]\n", err)
		return
	}

	var creatorResponse CreatorResponse
	if err := json.Unmarshal(bytes, &creatorResponse); err != nil {
		log.Printf("解析tiktok 查询用户信息响应失败: [%v]\n", err)
		return
	}

	if creatorResponse.Code != 0 {
		log.Printf("请求tiktok 查询用户信息失败: code: %d, msg:%s\n", creatorResponse.Code, creatorResponse.Message)
		return
	}

	if len(creatorResponse.Data.CreatorProfile) == 0 {
		log.Printf("tiktok 未找到用户: [%s]\n", uid)
		return
	}

	c := creatorResponse.Data.CreatorProfile[0]
	w.Write([]string{
		c.CreatorId,
		c.CreatorName,
		c.CreatorNickName,
		c.Region,
		strings.Join(c.ProductCategories, "|"),
		strconv.Itoa(c.FollowerCount),
		strconv.Itoa(c.VideoAvgViewCnt),
		c.CreatorOecuid,
		fmt.Sprintf("https://affiliate.tiktok.com/connection/creator/detail?cid=%s&pair_source=author_recommend", c.CreatorId),
	})
	w.Flush()
}

type XianzhiBatchResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
type DataList struct {
	UniqueID string `json:"unique_id"`
}
type Data struct {
	Count    int        `json:"count"`
	DataList []DataList `json:"dataList"`
}
