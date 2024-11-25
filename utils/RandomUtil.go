package utils

import "github.com/gox-studio/gona/uuid"

func GetRandomPassword() string {
	return uuid.New() //36
}
