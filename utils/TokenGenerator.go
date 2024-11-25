package utils

import (
	"strconv"
)

// GenerateToken 生成符合规则的令牌信息
func GenerateToken(productSecret string, uID int64) string {
	srcString := productSecret + strconv.Itoa(int(UID2SN(uID)))
	return GetMd5HexString(srcString)
	// return utils.GetRandomPassword()
}
