package actor

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

type Cooler struct {
	name string
	on   bool
	pin  rpio.Pin
}

func (actor Cooler) On() bool {
	actor.pin.High()
	fmt.Println("high")
	return true
}

func (actor Cooler) Off() bool {
	actor.pin.Low()
	fmt.Println("low")
	return true
}

func (actor Cooler) Init() bool {
	actor.pin = rpio.Pin(10)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return false
	}
	actor.pin.Output()
	return true
}

func (actor Cooler) TearDown() {
	rpio.Close()
}

func (actor Cooler) getStatus() bool {
	return true
}

func (actor Cooler) getType() ActorType {
	return COOLING
}
