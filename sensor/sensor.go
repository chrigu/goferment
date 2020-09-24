package sensor

type SensorUnit int

type SensorMsg struct {
	Value    float64
	SensorID string
	Unit     SensorUnit
}

const (
	TEMPERATURE SensorUnit = iota // always °C
)

type Sensor interface {
	GetUnit() SensorUnit
	GetValue() float64
	Init()
	StartCapture()
}
