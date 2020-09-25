package sensor

import (
	"fmt"
	"time"

	"github.com/yryz/ds18b20"
)

type Ds18b20 struct {
	value   float64
	sensors []string
}

func (sensor Ds18b20) GetValue() float64 {
	sensor.readSensor()
	fmt.Printf("for real temperature: %.2f°C\n", sensor.value)
	return sensor.value
}

func (sensor Ds18b20) GetUnit() SensorUnit {
	return TEMPERATURE
}

func (sensor Ds18b20) Init() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	sensor.sensors = sensors
	fmt.Printf("sensor IDs: %v\n", sensor.sensors)
}

func (sensor Ds18b20) readSensor() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	for _, i2cSensor := range sensors {
		t, err := ds18b20.Temperature(i2cSensor)
		if err == nil {
			fmt.Printf("sensor: %s temperature: %.2f°C\n", i2cSensor, t)
			sensor.value = t
			fmt.Printf("for real 1 temperature: %.2f°C\n", sensor.value)
		} else {
			fmt.Printf("error!")
		}
	}
}

func (sensor Ds18b20) StartCapture() {
	
	sensors, err := ds18b20.Sensors()

	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	go func() {
		for {
			sensor.readSensor()
			time.Sleep(time.Second * 10)
		}
	}()

}
