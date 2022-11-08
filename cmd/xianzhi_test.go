/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"testing"
)

func Test_fetchXianzhiNext(t *testing.T) {
	authorization = "ef017e58-0f8f-4c2b-be7a-eaa193b8c409"
	rsp, err := fetchXianzhiNext("GB", "1001", 1, 50)
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", rsp)

}

func Test_writeTikTokUserInfo(t *testing.T) {
	writeTikTokUserInfo(nil, "guinnessworldrecords")
}
