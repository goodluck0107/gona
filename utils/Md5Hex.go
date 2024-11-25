package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5HexString(srcString string) (result string) { //32
	h := md5.New()
	h.Write([]byte(srcString))
	result = hex.EncodeToString(h.Sum(nil))
	return
}
