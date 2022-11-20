package affiliate

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchCreators(t *testing.T) {

	// export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
	os.Setenv("https_proxy", "http://127.0.0.1:7890")
	curl := `curl 'https://affiliate.tiktok.com/api/v1/oec/affiliate/creator/marketplace/search?user_language=en&aid=4331&app_name=i18n_ecom_alliance&device_id=0&fp=verify_laesjkfh_TYWpYx37_7lvx_4u0P_8vHZ_3DJT4f1nyWmR&device_platform=web&cookie_enabled=true&screen_width=2560&screen_height=1440&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Mozilla&browser_version=5.0+(Macintosh%3B+Intel+Mac+OS+X+10_15_7)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F107.0.0.0+Safari%2F537.36+Edg%2F107.0.1418.42&browser_online=true&timezone_name=Asia%2FShanghai&shop_region=PH&msToken=HUoPdT1FjepGL7Yvez_hlheBYpXRYLzvsTfClF8HoI5kR4ZOmO4jZLK7O1SyKT3UPjlaeklOiT463X9sl33xel4LVqXweewiUMaeLNiNNCSfMMz_pCO7UWLz26in1YjJbmGV5qvH3to50MQJ&X-Bogus=DFSzswVuv9pw2D04S8ftdr7TlqCW&_signature=_02B4Z6wo00001NSsJNQAAIDDQHp.2xu3aZjUrCBAAFZha0' \
  -H 'authority: affiliate.tiktok.com' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
  -H 'content-type: application/json' \
  -H 'cookie: _fbp=fb.1.1668309682676.898292667; _ga=GA1.1.1683404244.1668309683; _ttp=2HTWil0nqAGJih5D9LYODPJbX9j; _tt_enable_cookie=1; passport_csrf_token=4a04be28f423912408f7a244e6748d8f; passport_csrf_token_default=4a04be28f423912408f7a244e6748d8f; sso_uid_tt_ads=14464b9dc754c8be786c930c8c870edfdfc77e59c0a6183db63502ae29a18184; sso_uid_tt_ss_ads=14464b9dc754c8be786c930c8c870edfdfc77e59c0a6183db63502ae29a18184; toutiao_sso_user_ads=a8fcf922b1be2281c43ba02d39409031; toutiao_sso_user_ss_ads=a8fcf922b1be2281c43ba02d39409031; sid_ucp_sso_v1_ads=1.0.0-KDYwMTcwYzMxNWNjZTIwYWE0OWZhZWI1ODQwYzAxNGVhODI5NzEzNmUKIAiCiMDiq4HrjGMQ0MXBmwYY5B8gDDDXgOiYBjgBQOsHEAEaA3NnMSIgYThmY2Y5MjJiMWJlMjI4MWM0M2JhMDJkMzk0MDkwMzE; ssid_ucp_sso_v1_ads=1.0.0-KDYwMTcwYzMxNWNjZTIwYWE0OWZhZWI1ODQwYzAxNGVhODI5NzEzNmUKIAiCiMDiq4HrjGMQ0MXBmwYY5B8gDDDXgOiYBjgBQOsHEAEaA3NnMSIgYThmY2Y5MjJiMWJlMjI4MWM0M2JhMDJkMzk0MDkwMzE; ttwid=1%7CKHXsJNOjNT_bMBtVnXZqyw9lQ7GI5vjvX0IiSc8af64%7C1668309715%7C686020d859f69ee0afd97de948f9d64b8ef3cbdd9a656ed3a246573bdd5670f1; SELLER_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcxNDA5Mjc4NDYyODE2NDMwMTAsIk9lY1VpZCI6NzQ5NDc0NTc0NjIyMTQwMjEyMiwiT2VjU2hvcElkIjo3NDk0NzQ1NzQ2MjIxNDAyMTIyLCJTaG9wUmVnaW9uIjoiIiwiR2xvYmFsU2VsbGVySWQiOjc0OTQ3NDU3NDYyMjE0MDIxMjIsIlNlbGxlcklkIjo3NDk0NzQ1NzQ2MjIxNDAyMTIyLCJleHAiOjE2NjgzOTYxMTksIm5iZiI6MTY2ODMwODcxOX0.rxKY9fkT0Iacg_Z5BA4XxzG3auDplmO-foJi5j7XCT0; SHOP_ID=7141191209187377414; s_v_web_id=verify_laesjkfh_TYWpYx37_7lvx_4u0P_8vHZ_3DJT4f1nyWmR; uid_tt_ads=14464b9dc754c8be786c930c8c870edfdfc77e59c0a6183db63502ae29a18184; uid_tt_ss_ads=14464b9dc754c8be786c930c8c870edfdfc77e59c0a6183db63502ae29a18184; sid_tt_ads=a8fcf922b1be2281c43ba02d39409031; sessionid_ads=a8fcf922b1be2281c43ba02d39409031; sessionid_ss_ads=a8fcf922b1be2281c43ba02d39409031; sid_guard_ads=a8fcf922b1be2281c43ba02d39409031%7C1668309740%7C863973%7CWed%2C+23-Nov-2022+03%3A21%3A53+GMT; sid_guard_tiktokshop=bd7f562bb3699bf8ed6949d01b2b1d84%7C1668309740%7C863973%7CWed%2C+23-Nov-2022+03%3A21%3A53+GMT; uid_tt_tiktokshop=2ffb251ce4db11072c56d3e93ae78d1fdfdc31024a6fe70228bfcaf0e3d6f2ff; uid_tt_ss_tiktokshop=2ffb251ce4db11072c56d3e93ae78d1fdfdc31024a6fe70228bfcaf0e3d6f2ff; sid_tt_tiktokshop=bd7f562bb3699bf8ed6949d01b2b1d84; sessionid_tiktokshop=bd7f562bb3699bf8ed6949d01b2b1d84; sessionid_ss_tiktokshop=bd7f562bb3699bf8ed6949d01b2b1d84; sid_ucp_v1_tiktokshop=1.0.0-KDc1NzI5YmU0YTA1NTI2YWZhMGI2MGExMzNjMTI3Yzc1Yjk4ZDg2YmYKGgiCiMDiq4HrjGMQ7MXBmwYY5B8gDDgBQOsHEAMaBm1hbGl2YSIgYmQ3ZjU2MmJiMzY5OWJmOGVkNjk0OWQwMWIyYjFkODQ; ssid_ucp_v1_tiktokshop=1.0.0-KDc1NzI5YmU0YTA1NTI2YWZhMGI2MGExMzNjMTI3Yzc1Yjk4ZDg2YmYKGgiCiMDiq4HrjGMQ7MXBmwYY5B8gDDgBQOsHEAMaBm1hbGl2YSIgYmQ3ZjU2MmJiMzY5OWJmOGVkNjk0OWQwMWIyYjFkODQ; i18next=en; csrf_session_id=a3be8efc75b58fe9e8fc91fdebdd3d24; _ga_BZBQ2QHQSP=GS1.1.1668322460.2.0.1668322460.0.0.0; msToken=HUoPdT1FjepGL7Yvez_hlheBYpXRYLzvsTfClF8HoI5kR4ZOmO4jZLK7O1SyKT3UPjlaeklOiT463X9sl33xel4LVqXweewiUMaeLNiNNCSfMMz_pCO7UWLz26in1YjJbmGV5qvH3to50MQJ; user_oec_info=0a532c88d2ba2842358aa3727b92a63bc9e282d2113ca7cf35ad36b174604673df578ac0010a4fd3674279361841a8845fd42c99fe29bdbb038463a50a4f67f6b5ce750da9d164f1b85c4f11e81233557b20f302671a490a3c51facb58743b1de2904f9aec4aa9dc93f7633f614c3cbf7ee97f8768aff90fdf40838c0386cb54244f14bceeb405358dc80c21d02da95cff6776591510a68ea10d1886d2f6f20d2201046647911a; msToken=oqHLUcBwuCEg62lq9NZblnDnJI9zLIFajphxr-PISEYzbXszfMa7FcxXUQ_LUPyPpMhgO2e38QNTi5R6VaBE6btzb1eJdbvsJD5DVhdX-fYMifQguhZUcQONTuEXm-rsCmzVXnd6VyQi_gQ5; odin_tt=0fff54938567f9bf833e26ba13e78a6d23113f76def71fd95ff986193474240542936607da63887a8cf0022f02a7fda2b287a468dcc34274f31eb25d127e81d4' \
  -H 'origin: https://affiliate.tiktok.com' \
  -H 'referer: https://affiliate.tiktok.com/connection/creator?enter_from=affiliate_find_creators' \
  -H 'sec-ch-ua: "Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.42' \
  --data-raw '{"request":{"follower_genders":[],"follower_age_groups":[],"managed_by_agency":[],"pagination":{"size":20,"page":0},"creator_score_range":[],"content_preference_range":[],"algorithm":3}}' \
  --compressed

	`

	var maxFollowerCount int = 100000
	rsp, err := searchCreators(curl, searchCreatorsRequest{
		Request: requestPayload{
			Algorithm:      3,
			FollowerCntMax: maxFollowerCount,
			Pagination: pagination{
				Size:           20,
				Page:           0,
			},
		},
	})

	assert.NoError(t, err)
	bs, _ := json.MarshalIndent(rsp, "", "\t")
	fmt.Println(string(bs))

}
