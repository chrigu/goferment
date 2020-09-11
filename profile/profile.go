package profile

import (
	"fmt"
	"goferment/hardware"
	"time"
)

type Step struct {
	temperature float32
	duration    int
	name        string
}

type Profile []Step

func profileLoop(ch chan string) {
	for {
		time.Sleep(time.Second)
		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		ch <- "tick"
	}
}

func commandLoop(ch, hwCh chan string) {
	for {
		message := <-ch
		fmt.Println("Message for cmdLoop: &v", message)
		hwCh <- message
	}
}

// StartProfile starts a defined temperature profile with one or multiple steps
func StartProfile(ch, cmdCh chan string) {

	// stepNumber := 0
	// startTime := time.Now().Unix()

	hwCh := make(chan string)

	go hardware.Hardware(hwCh)
	go profileLoop(ch)
	go commandLoop(cmdCh, hwCh)

}
