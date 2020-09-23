package profile

import (
	"fmt"
	"goferment/actor"
	"goferment/sensor"
	"time"
)

type Step struct {
	Temperature float32
	Duration    int
	Name        string
}

type Profile []Step

func profileLoop(ch chan string) {
	for {
		time.Sleep(time.Second)
		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		ch <- "tick"
	}
}

func commandLoop(ch chan string, sensorCh chan float64, actor actor.Actor) {
	for {
		select {
		case message := <-ch:
			fmt.Println("Message for actor: &v", message)
			if message == "on" {
				actor.On()
			} else {
				actor.Off()
			}
		case sensorData := <-sensorCh:
			fmt.Println("Temperature: &v", sensorData)
		}
	}
}

// StartProfile starts a defined temperature profile with one or multiple steps
func StartProfile(ch, cmdCh chan string) {

	// stepNumber := 0
	// startTime := time.Now().Unix()

	var ds18b20 sensor.Ds18b20
	var cooler actor.Cooler

	sensorCh := make(chan float64)

	ds18b20.Init(sensorCh)
	cooler.Init()

	go ds18b20.StartCapture()
	go profileLoop(ch)
	go commandLoop(cmdCh, sensorCh, cooler)

}
