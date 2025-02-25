package log

import (
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	l := &LogrusLogger{
		logger: logrus.New(),
	}
	conf := NewLogConf()

	l.Configure(conf)
	return l
}

func NewLogrusLoggerWithLevel(level LogLevel) *LogrusLogger {
	l := NewLogrusLogger()
	l.setLevel(level)

	return l
}

func (l *LogrusLogger) Configure(conf *LogConf) {
	// set log level
	l.setLevel(conf.Level)

	// set log format
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.FullTimestamp = true
	l.logger.SetFormatter(formatter)
	//l.logger.SetReportCaller(true)

	// set log rollback
	if conf.LogPath == "" {
		conf.LogPath = "./log/"
	}

	if conf.LogFile == "" {
		conf.LogFile = "app.log"
	}

	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, conf.LogFile),
		MaxSize:    conf.MaxSizePerFile,
		MaxBackups: conf.MaxBackups,
		MaxAge:     30,
		Compress:   conf.Compress,
	}
	l.logger.SetOutput(logFile)

}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusLogger) DebugF(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *LogrusLogger) InfoF(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *LogrusLogger) WarnF(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *LogrusLogger) ErrorF(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *LogrusLogger) setLevel(level LogLevel) {
	switch level {
	case DebugLevel:
		l.logger.Level = logrus.DebugLevel
	case InfoLevel:
		l.logger.Level = logrus.InfoLevel
	case WarnLevel:
		l.logger.Level = logrus.WarnLevel
	case ErrorLevel:
		l.logger.Level = logrus.ErrorLevel
	default:
		l.logger.Level = logrus.InfoLevel
	}
}
