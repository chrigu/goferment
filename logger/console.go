package logger

import (
	log "github.com/sirupsen/logrus"
)

type ConsoleLogger struct {
}

func (*ConsoleLogger) LogState(logEntry LogEntry) {
	log.Info(logEntry)
}
