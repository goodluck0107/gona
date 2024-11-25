package boost

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// String format struct or slice into json with indent
func String(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	return out.String()
}

// Printf is boost wrapper for fmt.Printf
func Printf(format string, args ...interface{}) {
	boostArgs := make([]interface{}, len(args))
	for i, arg := range args {
		boostArgs[i] = String(arg)
	}
	fmt.Printf(format, boostArgs...)
}

// Println is boost wrapper for fmt.Println
func Println(args ...interface{}) {
	boostArgs := make([]interface{}, len(args))
	for i, arg := range args {
		boostArgs[i] = String(arg)
	}
	fmt.Println(boostArgs...)
}

// Print is boost wrapper for fmt.Print
func Print(args ...interface{}) {
	boostArgs := make([]interface{}, len(args))
	for i, arg := range args {
		boostArgs[i] = String(arg)
	}
	fmt.Print(boostArgs...)
}
