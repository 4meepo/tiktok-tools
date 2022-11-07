/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "testing"

func Test_crawlCreators(t *testing.T) {
	curlFile = "../curl.txt"
	err := crawlCreators()
	if err != nil {
		t.Error(err)
	}
}
