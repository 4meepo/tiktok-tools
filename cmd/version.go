/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.2.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
