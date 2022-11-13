package affiliate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/4meepo/tiktok-tools/curl"
)

var ErrCurlCmd = errors.New("解析 curl 命令失败")

func searchCreators(curlSample string, request searchCreatorsRequest) (*searchCreatorsResponse, error) {
	// 利用curl中的 header和query param ...
	r, ok := curl.Parse(curlSample)
	if !ok {
		return nil, ErrCurlCmd
	}

	// 覆盖 request payload
	bytes, _ := json.Marshal(request)
	r.Body = string(bytes)

	req := r.AsHttpRequest()

	rsp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送searchCreators请求失败: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searchCreators请求失败: %s", rsp.Status)
	}

	rspBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取searchCreators 请求体失败: %w", err)
	}

	var response searchCreatorsResponse
	if err := json.Unmarshal(rspBytes, &response); err != nil {
		return nil, fmt.Errorf("解析searchCreators请求体失败: %w", err)
	}

	return &response, nil
}

// req

type searchCreatorsRequest struct {
	Request requestPayload `json:"request"`
}

type requestPayload struct {
	Algorithm              uint8         `json:"algorithm"`
	ContentPreferenceRange []interface{} `json:"content_preference_range,omitempty"`
	CreatorScoreRange      []interface{} `json:"creator_score_range,omitempty"`
	FollowerAgeGroups      []interface{} `json:"follower_age_groups,omitempty"`
	FollowerCntMax         *uint32       `json:"follower_cnt_max,omitempty"`
	FollowerCntMin         *uint32       `json:"follower_cnt_min,omitempty"`
	FollowerGenders        []interface{} `json:"follower_genders,omitempty"`
	Managed_by_agency      []interface{} `json:"managed_by_agency,omitempty"`
	Pagination             pagination    `json:"pagination,omitempty"`
}
type pagination struct {
	Size           int  `json:"size"` // from 0
	Page           int  `json:"page"`
	NextItemCursor *int `json:"next_item_cursor,omitempty"` // = size * page + 1, nil if = 0
}

// rsp

type searchCreatorsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		NextPagination struct {
			HasMore   bool `json:"has_more"`
			NextPage  int  `json:"next_page"`
			TotalPage int  `json:"total_page"`
			Total     int  `json:"total"`
		} `json:"next_pagination"`
		CreatorProfiles []*creatorProfile `json:"profiles"`
	} `json:"data"`
}
type creatorProfile struct {
	CreatorId         string   `json:"creator_id"`
	CreatorName       string   `json:"creator_name"`
	CreatorNickName   string   `json:"creator_nickname"`
	Region            string   `json:"region"`
	ProductCategories []string `json:"product_categories"`
	FollowerCount     uint32   `json:"follower_cnt"`
	VideoAvgViewCnt   uint32   `json:"video_avg_view_cnt"`
	VideoPubCnt       uint32   `json:"video_pub_cnt"`
	EcVideoAvgViewCnt uint32   `json:"ec_video_avg_view_cnt"`
	CreatorOecuid     string   `json:"creator_oecuid"`
	CreatorTtuid      string   `json:"creator_ttuid"`
}
