package utils

import "strings"

func GetInsertParam(length int) (result string) {
	result = "("
	if length >= 1 {
		result = result + "?"
	}
	if length > 1 {
		for i := 1; i < length; i = i + 1 {
			result = result + ",?"
		}
	}
	result = result + ")"
	return
}

func IsTxErr(err error) bool {
	if strings.Contains(err.Error(), "transaction") {
		return true
	}
	return false
}