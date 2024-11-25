package systemx

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/gox-studio/gona/channel/utils"
	"github.com/gox-studio/gona/logger"
)

// Function is to get current func name with package name
func Function() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// FileLine is to get caller's file name and line number
func FileLine() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("get file & line failed"))
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// SeriousPanic is to cause a fatal error when panic ocurrs.
func SeriousPanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprint(err, string(utils.Stack(3))))
			logger.Error(fmt.Sprint(err, string(debug.Stack())))
			Exit(1)
		}
	}()

	fn()
}

// SlightPanic is to cause no error when panic ocurrs.
func SlightPanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprint(err, string(utils.Stack(3))))
			logger.Error(fmt.Sprint(err, string(debug.Stack())))
		}
	}()

	fn()
}

// SilentPanic is to cause a normal error when panic ocurrs.
func SilentPanic(fn func() error) (e error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprint(err, string(utils.Stack(3))))
			logger.Error(fmt.Sprint(err, string(debug.Stack())))
		}
	}()

	return fn()
}
