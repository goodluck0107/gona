//go:build !windowsgui
// +build !windowsgui

package systemx

import (
	"os"
)

// Exit is a wrapper for os.Exit
func Exit(code int) {
	os.Exit(code)
}
