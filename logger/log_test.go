package logger_test

import (
	"testing"

	"gitee.com/andyxt/gona/logger"
)

func TestLog(t *testing.T) {
	logger.SetLogPath("/Users/cc/log/")
	logger.Info("test log")
}
