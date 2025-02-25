package log

import (
	"testing"
	"time"
)

func TestDebugLog(t *testing.T) {
	testLog(DebugLevel)
}

func TestInfoLog(t *testing.T) {
	testLog(InfoLevel)
}

func TestWarnLog(t *testing.T) {
	testLog(WarnLevel)
}

func TestErrorLog(t *testing.T) {
	testLog(ErrorLevel)
}

func testLog(level LogLevel) {
	conf := NewLogConf()
	conf.LogPath = "./"
	conf.LogFile = "app_" + level.String() + ".log"
	conf.MaxSizePerFile = 100
	conf.MaxBackups = 10
	conf.Compress = false

	conf.Level = level

	Configure(conf)
	Debug("this is debug log")
	Info("this is info log")
	Warn("this is warn log")
	Error("this error log")

	for i := 0; i < 100; i++ {
		if i%4 == 0 {
			DebugF("this is debug log, i = [%d]", i)
		}

		if i%4 == 1 {
			InfoF("this is info log, i = [%d]", i)
		}

		if i%4 == 2 {
			WarnF("this is warn log, i = [%d]", i)
		}

		if i%4 == 3 {
			ErrorF("this is error log, i = [%d]", i)
		}

		time.Sleep(time.Millisecond * 10)
	}
}
