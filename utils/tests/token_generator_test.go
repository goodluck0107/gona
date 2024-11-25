package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"gitee.com/andyxt/gona/utils"
)

func TestTokenGenerator(t *testing.T) {
	var token string = ""
	token = utils.GenerateToken("d994c3d5-c97b-4b1b-ab26-18fecd8ef636", 50362409)
	fmt.Println("d994c3d5-c97b-4b1b-ab26-18fecd8ef636 50362409 >", token, len(token))
	token = utils.GenerateToken("d994c3d5-c97b-4b1b-ab26-18fecd8ef636", utils.UID2SN(20023964))
	if token != "1fa0dee4c9e608bfc4484da5e8f0d746" {
		panic(errors.New("error:" + strconv.Itoa(int(20023964))))
	}
}
