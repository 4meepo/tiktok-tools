/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package xianzhi

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_fetchXianzhiNext(t *testing.T) {
	authorization := "c5256c44-a57c-479c-8219-6032f35cf792"
	userId := "8a669b48840f8e7501840f8e75a90000"
	rsp, err := queryByRegion("PH", userId, authorization, 343, 100)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.MarshalIndent(rsp, "", "  ")
	fmt.Println(string(bs))

}

func TestCrawlCreators(t *testing.T) {
	authorization := "c5256c44-a57c-479c-8219-6032f35cf792"
	userId := "8a669b48840f8e7501840f8e75a90000"
	CrawlCreators("PH", authorization, userId, 343, 50, time.Second)
}

func Test_randomDuration(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(randomDuration())
	}
}
