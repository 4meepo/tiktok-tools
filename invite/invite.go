package invite

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4meepo/tiktok-tools/curl"
	"github.com/sirupsen/logrus"
)

func Invite(curlString string) error {
	logrus.Info("邀约开始")
	r, ok := curl.Parse(curlString)
	if !ok {
		return fmt.Errorf("解析curl失败")
	}

	payload := constructPayload()

	// 输入用户的creators_id
	fmt.Printf("请输入creators_id:\n")
	input, err := bufio.NewReader(os.Stdin).ReadString(' ')
	if err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}

	creatrosID := strings.Split(strings.TrimSpace(input), "\n")
	logrus.Infof("准备邀约%d个用户", len(creatrosID))

	fmt.Printf("请确认是否开始邀约(y/n):")
	var confirm string
	fmt.Scanf("%s", &confirm)
	if confirm == "y" {

		for _, id := range creatrosID {
			time.Sleep(1 * time.Second)
			payload.CreatorID = id
			// 发送邀约请求
			bs, _ := json.Marshal(payload)
			r.Body = string(bs)
			req := r.AsHttpRequest()

			if err, remain := sendInvitation(req); err != nil {
				logrus.Errorf("邀约用户%s失败 ❌, %w", err)
				continue
			} else {
				logrus.Infof("邀约用户%s成功 ✅, 剩余可邀约数: %d", id, remain)
			}
		}
	}

	return nil
}

func trueOrFalse(s string) bool {
	if s == "y" {
		return true
	} else if s == "n" {
		return false
	} else {
		logrus.Error("输入错误, 请重新输入")
		os.Exit(1)
	}
	return false
}

func constructPayload() *invitationRequest {
	var payload = &invitationRequest{}

	// 样品
	fmt.Printf("是否提供免费试用? (y/n):")
	var sample string
	if _, err := fmt.Scanf("%s", &sample); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	payload.ProvidedFreeSample = trueOrFalse(sample)

	// 是否固定佣金
	fmt.Printf("是否固定佣金? (y/n):")
	var commission string
	if _, err := fmt.Scanf("%s", &commission); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if trueOrFalse(commission) {
		payload.IntentionalCooperation = append(payload.IntentionalCooperation, 1)
	}

	// 是否提成结佣
	fmt.Printf("是否提成结佣? (y/n):")
	var commission2 string
	if _, err := fmt.Scanf("%s", &commission2); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if trueOrFalse(commission2) {
		payload.IntentionalCooperation = append(payload.IntentionalCooperation, 2)
	}

	// 邀请消息
	fmt.Printf("请输入邀请消息:")
	var message string
	if _, err := fmt.Scanf("%s", &message); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if message == "" {
		logrus.Errorf("邀请消息不能为空")
		os.Exit(1)
	}
	payload.InvitationMessage = message

	// whatsapp country code
	fmt.Printf("请输入whatsapp国家代码(如MY#60):")
	var countryCode string
	if _, err := fmt.Scanf("%s", &countryCode); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if countryCode == "" {
		logrus.Errorf("whatsapp国家代码不能为空")
		os.Exit(1)
	}

	// whatsapp
	fmt.Printf("请输入whatsapp:")
	var whatsapp string
	if _, err := fmt.Scanf("%s", &whatsapp); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if whatsapp == "" {
		logrus.Errorf("whatsapp不能为空")
		os.Exit(1)
	}
	payload.ShopInvitationContactInfo = append(payload.ShopInvitationContactInfo, shopInvitationContactInfo{
		Field:       6,
		Title:       "",
		CountryCode: strings.TrimSpace(countryCode),
		Value:       strings.TrimSpace(whatsapp),
	})

	// 邮箱
	fmt.Printf("请输入邮箱:")
	var email string
	if _, err := fmt.Scanf("%s", &email); err != nil {
		logrus.Errorf("读取用户输入失败: %w", err)
		os.Exit(1)
	}
	if email == "" {
		logrus.Errorf("邮箱不能为空")
		os.Exit(1)
	}
	payload.ShopInvitationContactInfo = append(payload.ShopInvitationContactInfo, shopInvitationContactInfo{
		Field: 7,
		Value: strings.TrimSpace(email),
		Title: "",
	})
	return payload
}
