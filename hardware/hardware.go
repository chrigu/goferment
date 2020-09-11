package hardware

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

// StartProfile starts a defined temperature profile with one or multiple steps
func Hardware(ch chan string) {

	// stepNumber := 0
	// startTime := time.Now().Unix()

	var pin = rpio.Pin(10)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	pin.Output()

	for {
		cmd := <-ch

		switch {
		case cmd == "on":
			pin.High()
		case cmd == "off":
			pin.Low()
		}

	}

}
