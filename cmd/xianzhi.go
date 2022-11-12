/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

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
		return xianzhi.CrawlCreators(xianzhiRegion, authorization, xianzhiFromPage)
	},
}

var authorization, xianzhiRegion string
var xianzhiFromPage int

func init() {
	crawlCmd.AddCommand(xianzhiCmd)

	xianzhiCmd.Flags().StringVarP(&authorization, "authorization", "a", "", "先知网的 authorization")
	xianzhiCmd.Flags().StringVarP(&xianzhiRegion, "region", "r", "", "地区 VN--越南 MY--马来西亚 TH--泰国  PH--菲律宾")
	xianzhiCmd.Flags().IntVarP(&xianzhiFromPage, "fromPage", "p", 0, "从第几页开始爬取")
}
