package logger

import "log"

type Logger interface {
	Log(prefix string, msg string)
}

type LoggerBasic struct {
	appName    string
	loggerName string
}

func ProvideLogger(appName string) *LoggerBasic {
	return &LoggerBasic{
		appName:    appName,
		loggerName: "LoggerBasic",
	}
}

func (l *LoggerBasic) Log(prefix string, msg string) {
	log.Printf(
		"[%s] <logger %s>: %s:%s",
		l.appName, l.loggerName, prefix, msg,
	)
}
