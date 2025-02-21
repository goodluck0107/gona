package logger

import (
	"fmt"
)

func GetLogger() *logw {
	return &logw{}
}

func StartUp(v ...interface{}) {
	if printLevel >= LogLevelDebug {
		fmt.Println(v...)
	}
	startUpLogger.Println(v...)
}

func Debug(v ...interface{}) {
	if printLevel >= LogLevelDebug {
		fmt.Println(v...)
	}
	if logLevel < LogLevelDebug {
		return
	}
	debugLogger.Println(v...)
}

func Info(v ...interface{}) {
	if printLevel >= LogLevelInfo {
		fmt.Println(v...)
	}
	if logLevel < LogLevelInfo {
		return
	}
	infoLogger.Println(v...)
}

func Warn(v ...interface{}) {
	if printLevel >= LogLevelWarn {
		fmt.Println(v...)
	}
	if logLevel < LogLevelWarn {
		return
	}
	warnLogger.Println(v...)
}

func Error(v ...interface{}) {
	if printLevel >= LogLevelError {
		fmt.Println(v...)
	}
	if logLevel < LogLevelError {
		return
	}
	errorLogger.Println(v...)
}
