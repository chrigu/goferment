package profile

import (
	"testing"
)

func TestCheckTempertureTooLow(t *testing.T) {
	var temp float64 = 20
	step := ProfileStep{temp, 2 * 60, "Test"}
	tempComparision := step.checkTemperature(temp - 0.6)
	if tempComparision != TOO_COLD {
		t.Error("Expected TOO_COLD, got ", tempComparision)
	}
}

func TestCheckTempertureOk(t *testing.T) {
	var temp float64 = 20
	step := ProfileStep{temp, 2 * 60, "Test"}
	tempComparision := step.checkTemperature(temp - 0.1)
	if tempComparision != OK {
		t.Error("Expected OK, got ", tempComparision)
	}
}

func TestCheckTempertureTooHot(t *testing.T) {
	var temp float64 = 20
	step := ProfileStep{temp, 2 * 60, "Test"}
	tempComparision := step.checkTemperature(temp + 0.6)
	if tempComparision != TOO_HOT {
		t.Error("Expected TOO_HOT, got ", tempComparision)
	}
}
