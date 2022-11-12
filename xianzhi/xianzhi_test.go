/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package xianzhi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_fetchXianzhiNext(t *testing.T) {
	authorization := "036b4647-e3a0-4c5c-b1f0-2bef2505498a"
	rsp, err := queryByRegion("ID", authorization, 1, 100)
	if err != nil {
		t.Error(err)
	}

	bs, _ := json.MarshalIndent(rsp, "", "  ")
	fmt.Println(string(bs))

}
