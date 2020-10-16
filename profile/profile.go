package profile

import (
	"encoding/json"
	"fmt"
	"goferment/actor"
	"goferment/logger"
	"goferment/sensor"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type TempComparison int

const (
	LOWER_LIMIT TempComparison = iota
	OK
	UPPER_LIMIT
)

const LOOP_INTERVAL = 10

type ProfileStep struct {
	Temperature float64 `json:"temperature"`
	Duration    int     `json:"duration"`
	Name        string  `json:"name"`
}

type Profile struct {
	Steps []ProfileStep `json:"steps"`
}

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

func (currentStep *CurrentStep) activateStep() {

	if currentStep.active {
		return
	}

	currentStep.active = true
	currentStep.startTime = time.Now()
}

func (currentStep *CurrentStep) stepTimeLeft() float64 {
	duration := time.Since(currentStep.startTime)
	return float64(currentStep.Duration) - duration.Minutes()
}

func (currentStep *CurrentStep) hasEnded() bool {
	return currentStep.stepTimeLeft() < 0
}

func (currentStep *CurrentStep) coolerHysteresis(temperatureState TempComparison) bool {
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

func (currentStep *CurrentStep) heaterHysteresis(temperatureState TempComparison) bool {
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

func profileLoop(profile *Profile, ch chan string, sensor sensor.Sensor, heater, cooler actor.Actor) {

	consoleLogger := logger.ConsoleLogger{}
	dynamoLogger := logger.DynamoDbLogger{}

	hysterisisDelta := 1.0

	stepIndex := 0
	currentStep := CurrentStep{ProfileStep: &profile.Steps[stepIndex], active: false, delta: hysterisisDelta}
	for {

		// move to fn
		if currentStep.active && currentStep.hasEnded() {
			stepIndex++

			if stepIndex == len(profile.Steps) {
				break
			}

			currentStep = CurrentStep{ProfileStep: &profile.Steps[stepIndex], active: false, delta: hysterisisDelta}
		}

		// fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		sensorTemp := sensor.GetValue()
		fmt.Printf("Temperature: %v\n", sensorTemp)

		temperatureState := currentStep.checkTemperature(sensorTemp)

		if temperatureState == OK {
			currentStep.activateStep()
		}

		if cooler != nil {
			if currentStep.coolerHysteresis(temperatureState) {
				cooler.On()
			} else {
				cooler.Off()
			}
		}

		if heater != nil {
			if currentStep.heaterHysteresis(temperatureState) {
				heater.On()
			} else {
				heater.Off()
			}
		}

		// handle time elapsed in step

		// logEntry := logger.LogEntry{Datetime: time.Now(), Temperature: sensorTemp, TargetTemp: currentStep.Temperature, HeaterState: heater.GetStatus(), CoolerState: cooler.GetStatus()}
		logEntry := logger.LogEntry{Datetime: time.Now(), Temperature: sensorTemp, TargetTemp: currentStep.Temperature, HeaterState: heater.GetStatus(), CoolerState: false}
		log.Debugf("%v: %v %v", currentStep.Name, currentStep.active, currentStep.stepTimeLeft())

		consoleLogger.LogState(logEntry)
		dynamoLogger.LogState(logEntry)
		ch <- "tick"
		time.Sleep(LOOP_INTERVAL * time.Second)
	}
}

func commandLoop(ch chan string, actor actor.Actor) {
	for {
		select {
		/*
			commands
				actor on
				actor off
				temp +/-
				stop
				next step

			init loop
			log event
		*/
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

func ReadProfileFromFile(filename string) *Profile {
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened ", filename)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var profile Profile
	json.Unmarshal(byteValue, &profile)

	for i := 0; i < len(profile.Steps); i++ {
		fmt.Printf("step name: %v, duration: %v \n", profile.Steps[i].Name, profile.Steps[i].Duration)
	}

	return &profile
}

// StartProfile starts a defined temperature profile with one or multiple steps
func StartProfile(profile *Profile) (chan string, chan string) {

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
