/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"time"

	"github.com/4meepo/tiktok-tools/xianzhi"
	"github.com/spf13/cobra"
)

// xianzhiCmd represents the xianzhi command
var xianzhiCmd = &cobra.Command{
	Use:   "xianzhi",
	Short: "从 先知网 爬取达人数据",
	RunE: func(cmd *cobra.Command, args []string) error {
		if authorization == "" {
			return errors.New("authorization 不能为空")
		}
		if xianzhiRegion == "" {
			return errors.New("region 不能为空")
		}
		if xianzhiFromPage < 1 {
			return errors.New("fromPage 不能小于 1")
		}
		var d time.Duration
		var err error
		if d, err = time.ParseDuration(xianzhiDuration); err != nil {
			return errors.New("duration 格式错误")
		}
		return xianzhi.CrawlCreators(xianzhiRegion, authorization, xianzhiUserId, xianzhiFromPage, batchSize, d)
	},
}

var authorization, xianzhiUserId, xianzhiRegion string
var xianzhiFromPage, batchSize int
var xianzhiDuration string

func init() {
	crawlCmd.AddCommand(xianzhiCmd)

	xianzhiCmd.Flags().StringVarP(&authorization, "authorization", "a", "", "先知网的 authorization")
	xianzhiCmd.Flags().StringVarP(&xianzhiRegion, "region", "r", "", "地区 VN--越南 MY--马来西亚 TH--泰国  PH--菲律宾")
	xianzhiCmd.Flags().StringVarP(&xianzhiUserId, "userId", "u", "", "userId")
	xianzhiCmd.Flags().StringVarP(&xianzhiDuration, "duration", "d", "", "休息时间")
	xianzhiCmd.Flags().IntVarP(&xianzhiFromPage, "fromPage", "f", 0, "从第几页开始爬取")
	xianzhiCmd.Flags().IntVarP(&batchSize, "batchSize", "b", 0, "爬取批次大小")
	// todo 粉丝数范围
}
