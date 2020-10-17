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

type DynamoDbLogger struct {
	db        *dynamo.DB
	tableName string
}

func (dl *DynamoDbLogger) InitDb(region, awsProfile, tableName string) bool {
	dl.db = dynamo.New(session.New(), &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", awsProfile),
	})

	dl.tableName = tableName

	return false // todo: return state
}

func (dl *DynamoDbLogger) LogState(logEntry LogEntry) bool {
	fmt.Println(logEntry)

	table := dl.db.Table(dl.tableName)
	err := table.Put(dynamoLogEntryFromLogEntry(logEntry)).Run()

	fmt.Println(err)
	return true // todo: return state
}
