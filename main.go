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
	profileCh = make(chan string)
	profileCmdCh = make(chan string)

	server := server.CreateServer(webCh)
	go listener(webCh, profileCh, profileCmdCh)

	profile.StartProfile(profileCh, profileCmdCh)

	fmt.Println("Server starting")

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func listener(c1, c2, c3 chan string) {
	for i := 0; i < 100; i++ {
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
