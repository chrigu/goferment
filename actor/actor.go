package actor

type ActorType int

const (
	COOLING ActorType = iota
	HEATING
)

type Actor interface {
	On() bool
	Off() bool
	Init() bool // use err as return value
	TearDown()
	getStatus() bool
	getType() ActorType
}
