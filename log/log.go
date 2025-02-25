package log

import (
	"fmt"
	"path"
	"runtime"
	"sync"

	"github.com/jerryharbour/libgo/str"
)

var (
	loger Logger
	once  sync.Once
)

func Configure(conf *LogConf) {
	getLogger().Configure(conf)
}

func Debug(args ...interface{}) {
	getLogger().Debug(args...)
}

func Info(args ...interface{}) {
	getLogger().Info(args...)
}

func Warn(args ...interface{}) {
	getLogger().Warn(args...)
}

func Error(args ...interface{}) {
	getLogger().Error(args...)
}

func DebugF(format string, args ...interface{}) {
	getLogger().DebugF(addCaller(format), args...)
}

func InfoF(format string, args ...interface{}) {
	getLogger().InfoF(addCaller(format), args...)
}

func WarnF(format string, args ...interface{}) {
	getLogger().WarnF(addCaller(format), args...)
}

func ErrorF(format string, args ...interface{}) {
	getLogger().ErrorF(addCaller(format), args...)
}

func getLogger() Logger {
	once.Do(func() {
		loger = NewLogrusLogger()
	})

	return loger
}

func addCaller(format string) string {
	// call stack
	pc := make([]uintptr, 3)
	n := runtime.Callers(2, pc)
	if n <= 1 {
		return format
	}

	frame, ok := runtime.CallersFrames(pc[1:]).Next()
	if !ok {
		return format
	}

	return fmt.Sprintf("[%s:%s:%s]: %s", path.Base(frame.File), str.IntToString(frame.Line), str.TrimBfLastChar(frame.Function, "."), format)
}
