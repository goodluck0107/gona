package utils

import "gitee.com/andyxt/gona/uuid"

func GetRandomPassword() string {
	return uuid.New() //36
}
