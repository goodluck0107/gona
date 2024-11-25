package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"gona/utils"
)

func TestUIDGenerator(t *testing.T) {
	var uID int64 = 0
	uID = utils.SN2UID(50362409)
	fmt.Println("50362409 >", uID)
	fmt.Println("50362409 <", utils.UID2SN(uID))
	uID = utils.SN2UID(50362009)
	fmt.Println("50362009 >", uID)
	fmt.Println("50362009 <", utils.UID2SN(uID))

	for k := int64(0); k <= 80000000; k++ {
		uID := utils.SN2UID(k)
		if utils.UID2SN(uID) != k {
			panic(errors.New("back error:" + strconv.Itoa(int(k))))
		}
	}

}
