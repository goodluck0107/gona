//go:build windowsgui
// +build windowsgui

package systemx

import (
	"fmt"
	"os"
	"time"
)

// Exit is a wrapper for os.Exit
func Exit(code int) {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	count := 10
	for {
		select {
		case <-ticker.C:
			if count == 0 {
				os.Exit(code)
			} else {
				fmt.Printf("Exit in %d sec(s)\n", count)
			}
			count--
		}
	}
}
