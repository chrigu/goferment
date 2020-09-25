package profile

import (
	"fmt"
	"goferment/actor"
	"goferment/sensor"
	"time"
)

type TempComparison int

const (
	TOO_COLD TempComparison = iota
	OK
	TOO_HOT
)

type Step struct {
	Temperature float64
	Duration    int
	Name        string
}

func (step Step) checkTemperature(currentTemperature float64) TempComparison {
	if currentTemperature < step.Temperature-0.5 {
		return TOO_COLD
	} else if currentTemperature > step.Temperature+0.5 {
		return TOO_HOT
	} else {
		return OK
	}
}

type Profile []Step

func profileLoop(profile Profile, ch chan string, sensor sensor.Sensor, actor actor.Actor) {
	firstStep := profile[0]
	for {

		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		sensorTemp := sensor.GetValue()
		fmt.Printf("Temperature: %v\n", sensorTemp)

		temperatureState := firstStep.checkTemperature(sensorTemp)

		if temperatureState == TOO_COLD {
			actor.On()
		} else {
			actor.Off()
		}

		ch <- "tick"
		time.Sleep(5 * time.Second)
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
	step := Step{Temperature: 24, Duration: 2 * 60, Name: "Test"}
	profile := []Step{step}

	ds18b20 := &sensor.Ds18b20{}
	var cooler actor.Cooler

	ds18b20.Init()
	cooler.Init()

	ds18b20.StartCapture()
	go profileLoop(profile, ch, ds18b20, cooler)
	go commandLoop(cmdCh, cooler)

}
