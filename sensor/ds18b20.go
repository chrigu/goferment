package sensor

import (
	"fmt"
	"time"

	"github.com/yryz/ds18b20"
)

func Ds18b20(ch chan float64) {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	for {
		for _, sensor := range sensors {
			t, err := ds18b20.Temperature(sensor)
			if err == nil {
				fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
				ch <- t
			}
		}
		time.Sleep(time.Second * 20)
	}
}
