package logger

// https://github.com/guregu/dynamo

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type dynamoLogEntry struct {
	Datetime    time.Time `dynamo:"datetime"`
	HeaterState bool      `dynamo:"heaterState"`
	CoolerState bool      `dynamo:"coolerState"`
	Temperature float64   `dynamo:"temperature"`
	TargetTemp  float64   `dynamo:"targetTemp"`
	StepActive  bool      `dynamo:"stepActive"`
}

func dynamoLogEntryFromLogEntry(logEntry LogEntry) dynamoLogEntry {
	return dynamoLogEntry{Datetime: logEntry.Datetime, HeaterState: logEntry.HeaterState, CoolerState: logEntry.CoolerState,
		Temperature: logEntry.Temperature, StepActive: logEntry.StepActive}
}

type DynamoDbLogger struct{}

func (*DynamoDbLogger) LogState(logEntry LogEntry) {
	fmt.Println(logEntry)

	// todo: get from env
	db := dynamo.New(session.New(), &aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", "goferment"),
	})

	table := db.Table("LogEntries")

	err := table.Put(dynamoLogEntryFromLogEntry(logEntry)).Run()

	fmt.Println(err)
}
