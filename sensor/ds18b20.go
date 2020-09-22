package sensor

import (
	"fmt"
	"time"

	"github.com/yryz/ds18b20"
)

type Ds18b20 struct {
	value        float64
	sensors      []string
	valueChannel chan float64
}

func (sensor Ds18b20) GetValue() float64 {
	return sensor.value
}

func (sensor Ds18b20) GetUnit() SensorUnit {
	return TEMPERATURE
}

func (sensor Ds18b20) Init(valueChannel chan float64) {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	sensor.sensors = sensors
	sensor.valueChannel = valueChannel
}

func (sensor Ds18b20) StartCapture() {
	fmt.Printf("sensor IDs: %v\n", sensor.sensors)

	for {
		for _, i2cSensor := range sensor.sensors {
			t, err := ds18b20.Temperature(i2cSensor)
			if err == nil {
				fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
				sensor.valueChannel <- t
			}
		}
		time.Sleep(time.Second * 20)
	}
}
