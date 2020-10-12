package logger

import "time"

type Logger interface {
	LogState(LogEntry) bool
}

type LogEntry struct {
	Datetime    time.Time
	HeaterState bool
	CoolerState bool
	Temperature float64
	TargetTemp  float64
	StepActive  bool
}
