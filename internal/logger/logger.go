package logger

import "log"

type Logger interface {
	Log(prefix string, msg string)
}

type LoggerBasic struct {
	name       string
	loggerName string
}

type LoggerCount struct {
	name  string
	count int `wire:"-"` // Prevents Wire injection
}

func ProvideLoggerBasic(name string) *LoggerBasic {
	return &LoggerBasic{
		name:       name,
		loggerName: "LoggerBasic",
	}
}

func ProvideLoggerCount(name string) *LoggerCount {
	return &LoggerCount{
		name:  name,
		count: 0,
	}
}

func (l *LoggerBasic) Log(prefix string, msg string) {
	log.Printf(
		"[%s] <logger %s>: %s:%s",
		l.name, l.loggerName, prefix, msg,
	)
}

func (l *LoggerCount) Log(prefix string, msg string) {
	log.Printf(
		"[%s] <logger #%d>: %s:%s",
		l.name, l.count, prefix, msg,
	)

	l.count++
}
