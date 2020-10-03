package profile

import (
	"fmt"
	"goferment/actor"
	"goferment/sensor"
	"time"
)

type TempComparison int

const (
	LOWER_LIMIT TempComparison = iota
	OK
	UPPER_LIMIT
)

type ProfileStep struct {
	Temperature float64
	Duration    int
	Name        string
}

type Profile []ProfileStep

type CurrentStep struct {
	*ProfileStep
	active    bool
	startTime time.Time
	delta     float64
	coolerOn  bool
	heaterOn  bool
}

func (step *CurrentStep) checkTemperature(currentTemperature float64) TempComparison {
	if currentTemperature < step.Temperature-step.delta/2 {
		return LOWER_LIMIT
	} else if currentTemperature > step.Temperature+step.delta/2 {
		return UPPER_LIMIT
	} else {
		return OK
	}
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

func (currentStep CurrentStep) coolerHysteresis(temperatureState TempComparison) bool {
	switch temperatureState {
	case LOWER_LIMIT:
		currentStep.coolerOn = false
		return false
	case UPPER_LIMIT:
		currentStep.coolerOn = true
		return true
	default:
		return currentStep.coolerOn
	}
}

func (currentStep CurrentStep) heaterHysteresis(temperatureState TempComparison) bool {
	switch temperatureState {
	case LOWER_LIMIT:
		currentStep.heaterOn = true
		return true
	case UPPER_LIMIT:
		currentStep.heaterOn = false
		return false
	default:
		return currentStep.heaterOn
	}
}

func profileLoop(profile Profile, ch chan string, sensor sensor.Sensor, heater, cooler actor.Actor) {

	hysterisisDelta := 1.0

	currentStep := CurrentStep{ProfileStep: &profile[0], active: false, delta: hysterisisDelta}
	for {

		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		sensorTemp := sensor.GetValue()
		fmt.Printf("Temperature: %v\n", sensorTemp)

		temperatureState := currentStep.checkTemperature(sensorTemp)

		if temperatureState == OK {
			currentStep.activateStep(temperatureState)
		}

		if currentStep.coolerHysteresis(temperatureState) {
			cooler.On()
		} else {
			cooler.Off()
		}

		if currentStep.heaterHysteresis(temperatureState) {
			heater.On()
		} else {
			heater.Off()
		}

		// handle time elapsed in step

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
	// cooler := actor.NewTemperatureActor("cooling", actor.COOLING, 10)
	heater := actor.NewTemperatureActor("heater", actor.HEATING, 10)

	ds18b20.Init()

	ds18b20.StartCapture()
	go profileLoop(profile, ch, ds18b20, heater, nil)
	go commandLoop(cmdCh, heater)

	return cmdCh, ch

}
