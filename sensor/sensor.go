package sensor

type SensorUnit int

const (
	TEMPERATURE SensorUnit = iota // always °C
)

type Sensor interface {
	GetUnit() SensorUnit
	GetValue() float64
	Init(chan float64)
	StartCapture()
}
