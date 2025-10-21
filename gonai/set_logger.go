package gonai

import "github.com/goodluck0107/gona/internal/logger"

func SetLogger(l logger.Logger) {
	logger.Use(l)
}
