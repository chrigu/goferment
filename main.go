package main

import (
	"fmt"
	"log"

	"goferment/profile"
	"goferment/server"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

var webCh, profileCh, profileCmdCh chan string

func main() {

	webCh = make(chan string)
	var profileCh, profileCmdCh chan string

	step1 := profile.ProfileStep{Temperature: 24, Duration: 2 * 60, Name: "Test"}
	step2 := profile.ProfileStep{Temperature: 29, Duration: 2 * 60, Name: "Test"}

	fermentProfile := []profile.ProfileStep{step1, step2}

	server := server.CreateServer(webCh)
	go listener(webCh, profileCh, profileCmdCh)

	profileCmdCh, profileCh = profile.StartProfile(fermentProfile)

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
