package curl

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

type M map[string]interface{}

func TestParseTiktokCurl(t *testing.T) {
	request, _ := Parse(`curl 'https://api16-normal-useast1a.tiktokglobalshop.com/api/v1/oec/affiliate/creator/marketplace/recommendation?user_language=zh-CN&aid=6556&app_name=i18n_ecom_alliance&device_id=0&fp=verify_la662hr8_MGzIHxiS_W5KV_4QHp_BV4v_Bw4LQ7Pmh7yi&device_platform=web&cookie_enabled=true&screen_width=2048&screen_height=1280&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Mozilla&browser_version=5.0+(Macintosh%3B+Intel+Mac+OS+X+10_15_7)+AppleWebKit%2F537.36+(KHTML,+like+Gecko)+Chrome%2F107.0.0.0+Safari%2F537.36+Edg%2F107.0.1418.28&browser_online=true&timezone_name=Asia%2FShanghai&shop_region=GB&msToken=MH7GzFfseh5B9KKlcHWDRljgFizyOvwQGQce6vhGOfTpIf0U--b1CNMniEsA86QwSIT1vdtm9LNKg8WwWSnvAm3b75F3pJZ5Y6YVvKpWWRnWlvuCGX-H&X-Bogus=DFSzswVLZmaeeGaES0JLOOKMtak6&_signature=_02B4Z6wo00001uUuX9QAAIDDG2Ij5T03w9rlLltAANow90' \
  -H 'authority: api16-normal-useast1a.tiktokglobalshop.com' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
  -H 'content-type: application/json' \
  -H 'cookie: passport_csrf_token=c28e7da4dd75267c777f938ad8a8cf6e; passport_csrf_token_default=c28e7da4dd75267c777f938ad8a8cf6e; d_ticket=2e1e455b2127f5f377f356aa60a273a5c039d; sid_guard=745da5f251f52aa8002915358507bea1%7C1667787145%7C864000%7CThu%2C+17-Nov-2022+02%3A12%3A25+GMT; uid_tt=a4a479a62b3c50d7e4569702263c8449e5fed11b59414fe1f69e99e98412fbf8; uid_tt_ss=a4a479a62b3c50d7e4569702263c8449e5fed11b59414fe1f69e99e98412fbf8; sid_tt=745da5f251f52aa8002915358507bea1; sessionid=745da5f251f52aa8002915358507bea1; sessionid_ss=745da5f251f52aa8002915358507bea1; sid_ucp_v1=1.0.0-KGQ3MGRkOWEwNjllZTVlOGRhMWZkNmIyZmQ3ZTI0NzU2YjJlOTRkOGUKIAiBiIKU1Paqn2MQidOhmwYYnDMgDDDV-_uZBjgBQOsHEAEaA3NnMSIgNzQ1ZGE1ZjI1MWY1MmFhODAwMjkxNTM1ODUwN2JlYTE; ssid_ucp_v1=1.0.0-KGQ3MGRkOWEwNjllZTVlOGRhMWZkNmIyZmQ3ZTI0NzU2YjJlOTRkOGUKIAiBiIKU1Paqn2MQidOhmwYYnDMgDDDV-_uZBjgBQOsHEAEaA3NnMSIgNzQ1ZGE1ZjI1MWY1MmFhODAwMjkxNTM1ODUwN2JlYTE; store-idc=maliva; store-country-code=gb; store-country-code-src=uid; odin_tt=ea29aa1a1cecd22c321513b8b194e1fe81feedd07f0dfc74aa741cf611dd2d47f26097e7db1f44b52727685cca42f297b6e5d53bc2f8f926f132cdf8b696ecbd; ttwid=1%7CJamJ6DSQTUBINetCMp_9f7ZBH8J3PPob_yVRTm6_010%7C1667788419%7C5d051edf496f7cf3b7e0c800e669df2f6756cd16cb2ea86d86ec650947e15d8f; user_oec_info=0a53e4aba0bba95df89ca6898373eff72e727cf1806acc7c88e057f8649a00a133b5656a50e44f3a1cef58f82d54886765161648da168fbcaad7c3aa82232e45403ab2ac212b74eb34f18f61c31f972de5b8e9bbe91a490a3ca93e97f600d603385948df8fc9128f3410ff8c54cee1da5d4f9cbc3c190c240070e524bd4b9600ecf014bd91cc645f3ab9ef8780d5906594c45601ad10cbc8a00d1886d2f6f20d2201047ed3dfc1; msToken=XyoQ5kLQgijpbz1B4tgh2bW7zkWK0NaoWfZIwbUkJvhKNzV96HgcEu5O0f3B1eGR10Qt4miLh2jUF5P0CxFfo_25LuphWDKsb9gQceqQrUEyuyp9gGFr' \
  -H 'origin: https://affiliate.tiktokglobalshop.com' \
  -H 'referer: https://affiliate.tiktokglobalshop.com/' \
  -H 'sec-ch-ua: "Microsoft Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-site' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.28' \
  --data-raw '{"request":{"follower_genders":[],"follower_age_groups":[],"managed_by_agency":[],"pagination":{"size":20,"page":3,"next_item_cursor":61},"follower_cnt_max":5000,"follower_cnt_min":1000,"creator_score_range":[],"content_preference_range":[]}}' \
  --compressed

	`)

	fmt.Println(request.ToJson(true))

	x := request.AsHttpRequest()
	rsp, err := http.DefaultClient.Do(x)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rsp.Body.Close()
	fmt.Println(rsp.StatusCode)
}

func TestParse(t *testing.T) {

	addSample(t, "curl -XPUT http://api.sloths.com/sloth/4", M{
		"method": "PUT",
		"url":    "http://api.sloths.com/sloth/4",
	})

	addSample(t, "curl http://api.sloths.com", M{
		"method": "GET",
		"url":    "http://api.sloths.com",
	})

	addSample(t, "curl -H \"Accept-Encoding: gzip\" --compressed http://api.sloths.com", M{
		"method": "GET",
		"url":    "http://api.sloths.com",
		"header": M{
			"Accept-Encoding": "gzip",
		},
	})

	addSample(t, "curl -X DELETE http://api.sloths.com/sloth/4", M{
		"method": "DELETE",
		"url":    "http://api.sloths.com/sloth/4",
	})

	addSample(t, "curl -d \"foo=bar\" https://api.sloths.com", M{
		"method": "POST",
		"url":    "https://api.sloths.com",
		"header": M{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		"body": "foo=bar",
	})

	addSample(t, "curl -H \"Accept: text/plain\" --header \"User-Agent: slothy\" https://api.sloths.com", M{
		"method": "GET",
		"url":    "https://api.sloths.com",
		"header": M{
			"Accept":     "text/plain",
			"User-Agent": "slothy",
		},
	})

	addSample(t, "curl --cookie 'species=sloth;type=galactic' slothy https://api.sloths.com", M{
		"method": "GET",
		"url":    "https://api.sloths.com",
		"header": M{
			"Cookie": "species=sloth;type=galactic",
		},
	})

	addSample(t, "curl --location --request GET 'http://api.sloths.com/users?token=admin'", M{
		"method": "GET",
		"url":    "http://api.sloths.com/users?token=admin",
	})
}

func addSample(t *testing.T, url string, exp M) {
	request, _ := Parse(url)
	check(t, exp, request)
}

func check(t *testing.T, exp M, got *Request) {
	for key, value := range exp {
		switch key {
		case "method":
			assert.Equal(t, value, got.Method)
		case "url":
			assert.Equal(t, value, got.Url)
		case "body":
			assert.Equal(t, value, got.Body)
		case "header":
			headers := value.(M)
			for k, v := range headers {
				assert.Equal(t, v, got.Header[k])
			}
		}
	}
}
