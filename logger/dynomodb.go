package logger

// https://github.com/guregu/dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoDbLogger struct {
	url      string
	user     string
	password string
}

func (*DynamoDbLogger) LogState(logEntry LogEntry) {
	fmt.Println(logEntry)

	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})
	table := db.Table("LogEntries")

	err := table.Put(logEntry).Run()

	fmt.Println(err)
}
