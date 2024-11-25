package utils

import "gona/uuid"

func GetRandomPassword() string {
	return uuid.New() //36
}
