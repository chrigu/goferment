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

type ProfileStep struct {
	Temperature float64
	Duration    int
	Name        string
}

func (step *ProfileStep) checkTemperature(currentTemperature float64) TempComparison {
	if currentTemperature < step.Temperature-0.5 {
		return TOO_COLD
	} else if currentTemperature > step.Temperature+0.5 {
		return TOO_HOT
	} else {
		return OK
	}
}

type Profile []ProfileStep

type CurrentStep struct {
	*ProfileStep
	active    bool
	startTime time.Time
}

func (currentStep CurrentStep) activateStep(temperatureState TempComparison) {

	if currentStep.active {
		return
	}

	if temperatureState == OK {
		currentStep.active = true
		currentStep.startTime = time.Now()
	}
}

func (currentStep CurrentStep) stepTimeLeft() {

}

func profileLoop(profile Profile, ch chan string, sensor sensor.Sensor, heater, cooler actor.Actor) {
	currentStep := CurrentStep{ProfileStep: &profile[0], active: false}
	for {

		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		sensorTemp := sensor.GetValue()
		fmt.Printf("Temperature: %v\n", sensorTemp)

		temperatureState := currentStep.checkTemperature(sensorTemp)
		currentStep.activateStep(temperatureState)

		switch temperatureState {
		case TOO_COLD:
			heater.On()
		case TOO_HOT:
			cooler.On()
		default:
			heater.Off()
			cooler.Off()
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
func StartProfile(profile Profile) (chan string, chan string) {

	cmdCh := make(chan string)
	ch := make(chan string)

	// stepNumber := 0
	// startTime := time.Now().Unix()

	ds18b20 := &sensor.Ds18b20{}
	cooler := actor.NewTemperatureActor("cooling", actor.COOLING, 10)

	ds18b20.Init()
	cooler.Init()

	ds18b20.StartCapture()
	go profileLoop(profile, ch, ds18b20, nil, cooler)
	go commandLoop(cmdCh, cooler)

	return cmdCh, ch

}
