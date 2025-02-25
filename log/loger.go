package log

import (
	"io"
)

type Logger interface {
	Configure(conf *LogConf)
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	DebugF(format string, args ...interface{})
	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})
}

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "unknow"
	}
}

type LogConf struct {
	Level          LogLevel
	Format         string
	Output         io.Writer
	LogPath        string
	LogFile        string
	MaxSizePerFile int
	MaxBackups     int
	Compress       bool
}

func NewLogConf() *LogConf {
	return &LogConf{
		Level:          InfoLevel,
		Format:         "",
		Output:         nil,
		LogPath:        "./log/",
		LogFile:        "app.log",
		MaxSizePerFile: 100,
		MaxBackups:     10,
		Compress:       true,
	}
}
