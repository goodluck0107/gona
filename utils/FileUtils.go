package utils

import "os"

func IsFileExist(fullPath string) (isExist bool) {
	isExist = false
	if _, err := os.Stat(fullPath); err == nil {
		isExist = true
	}
	return
}