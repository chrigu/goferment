package main

import (
	"fmt"
	"log"
	"os"

	"goferment/logger"
	"goferment/profile"
	"goferment/server"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

var webCh, profileCh, profileCmdCh chan string

func test() {

	// some := logger.LogEntry{Datetime: time.Now(), HeaterState: true, Temperature: 24.5}

}

func main() {

	test()

	awsRegion := os.Getenv("AWS_REGION")
	awsProfile := os.Getenv("AWS_PROFILE")
	awsTableName := os.Getenv("AWS_TABLENAME")

	consoleLogger := &logger.ConsoleLogger{}
	dynamoLogger := &logger.DynamoDbLogger{}
	dynamoLogger.InitDb(awsRegion, awsProfile, awsTableName)

	webCh = make(chan string)
	var profileCh, profileCmdCh chan string

	fermentProfile := profile.ReadProfileFromFile("profile/test-profile.json")

	server := server.CreateServer(webCh)
	profileCmdCh, profileCh = profile.StartProfile(fermentProfile, []logger.Logger{consoleLogger, dynamoLogger})

	go listener(webCh, profileCh, profileCmdCh)

	fmt.Println("Server starting")

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func listener(c1, c2, c3 chan string) {
	for {
		select {
		case webMsg := <-c1:
			fmt.Println("received", webMsg)
			c3 <- webMsg
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			if msg2 == "off" {
				c3 <- "off"
			}
		}
	}
}
