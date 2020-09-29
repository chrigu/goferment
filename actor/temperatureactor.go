package actor

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

type TemperatureActor struct {
	Name          string
	on            bool
	pin           rpio.Pin
	RegulatorType ActorType
}

func (actor *TemperatureActor) On() bool {
	actor.pin = actor.initPin(10)
	actor.pin.High()
	fmt.Println("high")
	return true
}

func (actor *TemperatureActor) Off() bool {
	actor.pin = actor.initPin(10)
	actor.pin.Low()
	fmt.Println("low")
	return true
}

func (actor *TemperatureActor) Init() bool {
	actor.pin = actor.initPin(10)
	return true
}

func (actor *TemperatureActor) initPin(pinId int) rpio.Pin {
	pin := rpio.Pin(pinId)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return pin
	}
	pin.Output()
	return pin
}

func (actor *TemperatureActor) TearDown() {
	rpio.Close()
}

func (actor *TemperatureActor) getStatus() bool {
	return true
}

func (actor *TemperatureActor) getType() ActorType {
	return COOLING
}

func NewTemperatureActor(name string, actorType ActorType, pinId int) *TemperatureActor {
	pin := rpio.Pin(pinId)
	tr := TemperatureActor{Name: name, pin: pin, RegulatorType: actorType}
	return &tr
}
