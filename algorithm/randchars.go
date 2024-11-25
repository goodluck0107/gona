package algorithm

import (
	"bytes"
	"strings"
)

// 取得大写的随机字母
func RandUpString(l int) string {
	var result bytes.Buffer
	var temp byte
	for i := 0; i < l; {
		if randInt(65, 91) != temp {
			temp = randInt(65, 91)
			result.WriteByte(temp)
			i++
		}
	}
	return result.String()
}

// 取得小写的随机字母
func RandLowString(l int) string {
	return strings.ToLower(RandUpString(l))
}

// 取得数字随机个数
func RandIntString(l int) string {
	var result bytes.Buffer
	var temp byte
	for i := 0; i < l; {
		if randInt(48, 57) != temp {
			temp = randInt(48, 57)
			result.WriteByte(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) byte {
	return byte(RandIntn(min, max))
}
