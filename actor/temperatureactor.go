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
	fmt.Println("on1")
	fmt.Println("on", actor.pin)
	actor.pin.High()
	fmt.Println("high")
	return true
}

func (actor *TemperatureActor) Off() bool {
	actor.pin.Low()
	fmt.Println("low")
	return true
}

func (actor *TemperatureActor) Init() bool {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return false
	}
	actor.pin.Output()
	return true
}

func (actor *TemperatureActor) TearDown() {
	rpio.Close()
}

func (actor *TemperatureActor) GetStatus() bool {
	return true
}

func (actor *TemperatureActor) GetType() ActorType {
	return COOLING
}

func NewTemperatureActor(name string, actorType ActorType, pinId int) *TemperatureActor {
	pin := rpio.Pin(pinId)
	tr := TemperatureActor{Name: name, pin: pin, RegulatorType: actorType}

	tr.Init()
	return &tr
}
