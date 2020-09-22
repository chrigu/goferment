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

func commandLoop(ch, hwCh chan string, sensorCh chan float64) {
	for {
		select {
		case message := <-ch:
			fmt.Println("Message for cmdLoop: &v", message)
			hwCh <- message
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

	hwCh := make(chan string)
	sensorCh := make(chan float64)

	ds18b20.Init(sensorCh)

	go actor.Hardware(hwCh)
	go ds18b20.StartCapture()
	go profileLoop(ch)
	go commandLoop(cmdCh, hwCh, sensorCh)

}
