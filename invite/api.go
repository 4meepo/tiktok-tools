package invite

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}

func sendInvitation(req *http.Request) (err error, remain int) {
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("邀约失败, http状态码: %d", rsp.StatusCode), remain
	}

	bs, err := io.ReadAll(rsp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %w", err), remain
	}

	var rspBody struct {
		Code   int    `json:"code"`
		Msg    string `json:"msg"`
		Remain int    `json:"invited_remain_count"`
	}
	err = json.Unmarshal(bs, &rspBody)
	if err != nil {
		return fmt.Errorf("解析响应体失败: %w", err), remain
	}

	if rspBody.Code != 0 {
		return fmt.Errorf("邀约失败, code: %d, msg: %s", rspBody.Code, rspBody.Msg), remain
	}

	remain = rspBody.Remain

	return nil, remain
}

type invitationRequest struct {
	IntentionalCooperation    []int                       `json:"intentional_cooperation"`
	ProvidedFreeSample        bool                        `json:"provided_free_sample"`
	InvitationMessage         string                      `json:"invitation_message"`
	ShopInvitationContactInfo []shopInvitationContactInfo `json:"shop_invitation_contact_info"`
	CreatorID                 string                      `json:"creator_id"`
}
type shopInvitationContactInfo struct {
	Value       string `json:"value"`
	Field       int    `json:"field"`
	Title       string `json:"title"`
	CountryCode string `json:"country_code,omitempty"`
}
