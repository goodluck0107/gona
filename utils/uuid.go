package utils

import "github.com/gofrs/uuid"

func UUID() string {
	return uuid.Must(uuid.NewV4()).String() //36
}
