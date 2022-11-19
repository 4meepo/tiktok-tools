/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/4meepo/tiktok-tools/affiliate"
	"github.com/spf13/cobra"
)

// affiliateCmd represents the affiliate command
var affiliateCmd = &cobra.Command{
	Use:   "affiliate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 读取 curl 命令
		f, err := os.Open(curlFile)
		if err != nil {
			return fmt.Errorf("打开文件%s失败: %w", curlFile, err)
		}
		defer f.Close()
		bytes, err := io.ReadAll(f)
		if err != nil {
			return fmt.Errorf("读取字节流失败: %w", err)
		}

		// 解析 中场 休息时间
		d, err := time.ParseDuration(affiliateSleepDuration)
		if err != nil {
			return err
		}

		if affiliateRegion == "" {
			return errors.New("region 不能为空")
		}

		return affiliate.CrawlAffiliateCreators(mysqlHost, string(bytes), affiliateRegion, d, affiliatePageSize, affiliateMaxBatch)
	},
}

var affiliateRegion, affiliateSleepDuration string
var affiliatePageSize, affiliateMaxBatch int
var mysqlHost string

func init() {
	crawlCmd.AddCommand(affiliateCmd)
	affiliateCmd.Flags().StringVarP(&mysqlHost, "mysqlhost", "", "ecs", "mysql host")
	affiliateCmd.Flags().StringVarP(&affiliateRegion, "region", "r", "", "curl 命令所处的文件")
	affiliateCmd.Flags().StringVarP(&affiliateSleepDuration, "duration", "d", "", "每爬取1000条数据后休息的时间")
	affiliateCmd.Flags().IntVarP(&affiliatePageSize, "pageSize", "p", 20, "每页的数据量")
	affiliateCmd.Flags().IntVarP(&affiliateMaxBatch, "maxBatch", "m", 1000, "每页的数据量")

}
