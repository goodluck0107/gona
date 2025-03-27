package gonai

import "gitee.com/andyxt/gona/internal/logger"

func SetLogger(l logger.Logger) {
	logger.Use(l)
}
