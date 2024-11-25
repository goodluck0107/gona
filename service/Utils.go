package service

import (
	"unicode"
	"unicode/utf8"
)

func isExported(name string) bool {
	w, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(w)
}
