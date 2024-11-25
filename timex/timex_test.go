package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestTimex(t *testing.T) {
	configJson := "{\"zone\": \"UTC-3\",\"delta\": 3600,\"fake\": 0}"
	Init(configJson)
	fmt.Println("timex.Now():", Now())
	fmt.Println("timex.Now().Unix():", Now().Unix())
	fmt.Println("time.Now():", time.Now())
	fmt.Println("time.Now().Unix():", time.Now().Unix())
}
