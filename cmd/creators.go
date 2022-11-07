/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/4meepo/tiktok-tools/curl"
	"github.com/spf13/cobra"
)

// creatorsCmd represents the creators command
var creatorsCmd = &cobra.Command{
	Use:   "creators",
	Short: "爬取达人信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		return crawlCreators()
	},
}

var curlFile, outputfile string

func init() {
	crawlCmd.AddCommand(creatorsCmd)

	creatorsCmd.Flags().StringVarP(&curlFile, "file", "f", "", "curl 命令所处的文件")
	creatorsCmd.Flags().StringVarP(&outputfile, "output", "o", "", "输出文件名")
}

func crawlCreators() error {
	f, err := os.Open(curlFile)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("读取字节流失败: %w", err)
	}

	r, ok := curl.Parse(string(bytes))
	if !ok {
		return errors.New("解析 curl 命令失败")
	}

	var payload CreatorPayload
	err = json.Unmarshal([]byte(r.Body), &payload)
	if err != nil {
		return fmt.Errorf("解析creator请求体失败: %w", err)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	go elegantShutdown(cancelFn)

	var creators []CreatorProfile

	// 循环读取数据, 直至异常
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			saveFile(creators)
			log.Println("程序退出...")
			return nil
		default:
			// 1. 构造请求
			p := payload
			size := 20
			p.Request.Pagination.Size = size
			p.Request.Pagination.Page = i
			p.Request.Pagination.NextItemCursor = (size * i) + 1

			payloadBytes, _ := json.Marshal(p)
			r.Body = string(payloadBytes)
			httpRequest := r.AsHttpRequest()
			// 2. 发送请求
			httpResponse, err := httpClient.Do(httpRequest)
			if err != nil {
				return fmt.Errorf("发送http请求失败: %w", err)
			}
			httpResponseBytes, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				return fmt.Errorf("读取http响应失败: %w", err)
			}
			var response CreatorResponse
			if err := json.Unmarshal(httpResponseBytes, &response); err != nil {
				return fmt.Errorf("解析http响应失败: %w", err)
			}
			if response.Code != 0 {
				return fmt.Errorf("http响应异常: %d %s", response.Code, response.Message)
			}
			creators = append(creators, response.Data.CreatorProfile...)
			log.Printf("已爬取%d个达人", len(creators))
		}
	}
}

func saveFile(creators []CreatorProfile) {
	log.Println("保存文件至csv...")
	f, err := os.Create(fmt.Sprintf("%s.csv", strings.TrimSpace(outputfile)))
	if err != nil {
		log.Println("创建文件失败: ", err)
		return
	}
	defer f.Close()

	// 写入csv
	w := csv.NewWriter(f)
	w.Write([]string{"creator_id", "creator_name", "creator_nickname", "region", "product_categories", "follower_cnt", "creator_oecuid", "主页"})
	for _, c := range creators {
		w.Write([]string{
			c.CreatorId,
			c.CreatorName,
			c.CreatorNickName,
			c.Region,
			strings.Join(c.ProductCategories, "|"),
			strconv.Itoa(c.FollowerCount),
			c.CreatorOecuid,
			fmt.Sprintf("https://affiliate.tiktokglobalshop.com/connection/creator/detail?cid=%s&pair_source=author_recommend&enter_from=affiliate_home_page", c.CreatorId),
		})
	}
	w.Flush()
}

type CreatorPayload struct {
	Request Request `json:"request"`
}
type Pagination struct {
	Size           int `json:"size"`
	Page           int `json:"page"`
	NextItemCursor int `json:"next_item_cursor"`
}
type Request struct {
	FollowerGenders        []interface{} `json:"follower_genders"`
	FollowerAgeGroups      []interface{} `json:"follower_age_groups"`
	ManagedByAgency        []interface{} `json:"managed_by_agency"`
	Pagination             Pagination    `json:"pagination"`
	FollowerCntMax         int           `json:"follower_cnt_max"`
	FollowerCntMin         int           `json:"follower_cnt_min"`
	CreatorScoreRange      []interface{} `json:"creator_score_range"`
	ContentPreferenceRange []interface{} `json:"content_preference_range"`
}

type CreatorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    struct {
		CreatorProfile []CreatorProfile `json:"profiles"`
	} `json:"data"`
}
type CreatorProfile struct {
	CreatorId         string   `json:"creator_id"`
	CreatorName       string   `json:"creator_name"`
	CreatorNickName   string   `json:"creator_nickname"`
	Region            string   `json:"region"`
	ProductCategories []string `json:"product_categories"`
	FollowerCount     int      `json:"follower_cnt"`
	CreatorOecuid     string   `json:"creator_oecuid"`
}

// 优雅停机
func elegantShutdown(cancelFn context.CancelFunc) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
	cancelFn()
}
