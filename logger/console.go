package logger

import (
	"fmt"
)

type ConsoleLogger struct {
}

func (*ConsoleLogger) LogState(logEntry LogEntry) {
	fmt.Println(logEntry)
}
