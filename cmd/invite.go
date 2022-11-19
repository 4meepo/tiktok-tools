/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/4meepo/tiktok-tools/invite"
	"github.com/spf13/cobra"
)

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "达人邀约",

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

		return invite.Invite(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(inviteCmd)
}
