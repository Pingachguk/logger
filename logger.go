package logger

import (
	"fmt"
	"log"
	"os"
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

type LogLevel int

const (
	ErrorLevel LogLevel = iota
	WarningLevel
	InfoLevel
)

func NewLogExtended() LogExtended {
	return LogExtended{
		Logger:   log.New(os.Stderr, "", log.Ldate),
		logLevel: ErrorLevel,
	}
}

func (l *LogExtended) SetLogLevel(level LogLevel) error {
	switch level {
	case ErrorLevel, WarningLevel, InfoLevel:
		l.logLevel = level
	default:
		return fmt.Errorf("undefiend level: %d", level)
	}

	return nil
}

func (l LogExtended) Infoln(msg string) {
	l.println(InfoLevel, "[Info]", msg)
}

func (l LogExtended) Warnln(msg string) {
	l.println(WarningLevel, "[Warning]", msg)
}

func (l LogExtended) Errorln(msg string) {
	l.println(ErrorLevel, "[Error]", msg)
}

func (l LogExtended) println(level LogLevel, prefix, msg string) {
	if l.logLevel < level {
		return
	}

	l.Logger.Printf("%s %s", prefix, msg)
}
