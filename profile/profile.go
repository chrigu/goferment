package profile

import (
	"fmt"
	"goferment/actor"
	"goferment/sensor"
	"time"
)

type Step struct {
	Temperature float64
	Duration    int
	Name        string
}

type Profile []Step

func profileLoop(ch chan string, sensor sensor.Sensor) {
	for {
		time.Sleep(5 * time.Second)

		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		fmt.Printf("Temperature: %v\n", sensor.GetValue())
		ch <- "tick"
	}
}

func commandLoop(ch chan string, actor actor.Actor) {
	for {
		select {
		case message := <-ch:
			fmt.Println("Message for actor: &v", message)
			if message == "on" {
				actor.On()
			} else {
				actor.Off()
			}
		}
	}
}

// StartProfile starts a defined temperature profile with one or multiple steps
func StartProfile(ch, cmdCh chan string) {

	// stepNumber := 0
	// startTime := time.Now().Unix()

	var ds18b20 sensor.Ds18b20
	var cooler actor.Cooler

	ds18b20.Init()
	cooler.Init()

	ds18b20.StartCapture()
	go profileLoop(ch, ds18b20)
	go commandLoop(cmdCh, cooler)

}
