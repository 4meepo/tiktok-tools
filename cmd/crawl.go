/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "爬取",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crawl called")
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
}
