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
	Short: "爬取tiktok 达人信息",
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

		if affiliateFollowerFrom <= 0 {
			return errors.New("followerFrom 不能小于等于0")
		}

		return affiliate.CrawlAffiliateCreators(mysqlHost, string(bytes), d, affiliateFollowerFrom, affiliatePageSize, threshold, interval)
	},
}

var affiliateSleepDuration string
var affiliatePageSize, threshold int
var mysqlHost string
var affiliateFollowerFrom int // 从多少粉丝开始爬取, 间隔为100
var interval int              // api 爬取间隔

func init() {
	crawlCmd.AddCommand(affiliateCmd)
	affiliateCmd.Flags().StringVarP(&mysqlHost, "mysqlhost", "", "ecs", "mysql host")
	affiliateCmd.Flags().StringVarP(&affiliateSleepDuration, "duration", "d", "", "连续发送多少次请求后休息一段时间")
	affiliateCmd.Flags().IntVarP(&interval, "interval", "i", 30, "api请求平均间隔")
	affiliateCmd.Flags().IntVarP(&affiliatePageSize, "pageSize", "", 20, "每页的数据量")
	affiliateCmd.Flags().IntVarP(&threshold, "threshold", "", 50, "连续发起多少次请求后休息")
	affiliateCmd.Flags().IntVarP(&affiliateFollowerFrom, "followerFrom", "", 0, "从多少粉丝开始爬取, 间隔为100")

}
