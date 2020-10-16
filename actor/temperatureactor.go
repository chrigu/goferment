package actor

import (
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

type TemperatureActor struct {
	Name          string
	on            bool
	pin           rpio.Pin
	RegulatorType ActorType
}

func (actor *TemperatureActor) On() bool {
	log.Debugf("Pin %v high", actor.pin)
	actor.pin.High()
	return true
}

func (actor *TemperatureActor) Off() bool {
	log.Debugf("Pin %v low", actor.pin)
	actor.pin.Low()
	return true
}

func (actor *TemperatureActor) Init() bool {
	if err := rpio.Open(); err != nil {
		log.Warnf("Pin %v could not be intialized!", actor.pin)
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
