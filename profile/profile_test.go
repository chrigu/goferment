package profile

import (
	"testing"
)

func TestCheckTempertureTooLow(t *testing.T) {
	var temp float64 = 20
	profileStep := ProfileStep{temp, 2 * 60, "Test"}
	currentStep := CurrentStep{ProfileStep: &profileStep, delta: 1.0}
	tempComparision := currentStep.checkTemperature(temp - 0.6)
	if tempComparision != LOWER_LIMIT {
		t.Error("Expected TOO_COLD, got ", tempComparision)
	}
}

func TestCheckTempertureOk(t *testing.T) {
	var temp float64 = 20
	profileStep := ProfileStep{temp, 2 * 60, "Test"}
	currentStep := CurrentStep{ProfileStep: &profileStep, delta: 1.0}
	tempComparision := currentStep.checkTemperature(temp - 0.1)
	if tempComparision != OK {
		t.Error("Expected OK, got ", tempComparision)
	}
}

func TestCheckTempertureTooHot(t *testing.T) {
	var temp float64 = 20
	profileStep := ProfileStep{temp, 2 * 60, "Test"}
	currentStep := CurrentStep{ProfileStep: &profileStep, delta: 1.0}
	tempComparision := currentStep.checkTemperature(temp + 0.6)
	if tempComparision != UPPER_LIMIT {
		t.Error("Expected TOO_HOT, got ", tempComparision)
	}
}

func TestCheckCoolerOff(t *testing.T) {
	tempComp := LOWER_LIMIT
	currentStep := CurrentStep{coolerOn: true}
	coolerOn := currentStep.coolerHysteresis(tempComp)
	if coolerOn || currentStep.coolerOn {
		t.Error("Cooler should be off")
	}
}

func TestCheckCoolerOffWhileOk(t *testing.T) {
	tempComp := OK
	currentStep := CurrentStep{coolerOn: true}
	coolerOn := currentStep.coolerHysteresis(tempComp)
	if !coolerOn || !currentStep.coolerOn {
		t.Error("Cooler should be on")
	}
}

func TestCheckCoolerOnWhileOk(t *testing.T) {
	tempComp := OK
	currentStep := CurrentStep{coolerOn: false}
	coolerOn := currentStep.coolerHysteresis(tempComp)
	if coolerOn || currentStep.coolerOn {
		t.Error("Cooler should be off")
	}
}

func TestCheckHeaterOff(t *testing.T) {
	tempComp := LOWER_LIMIT
	currentStep := CurrentStep{heaterOn: true}
	heaterOn := currentStep.heaterHysteresis(tempComp)
	if !heaterOn || !currentStep.heaterOn {
		t.Error("Heater should be on")
	}
}

func TestCheckHeaterOffWhileOk(t *testing.T) {
	tempComp := OK
	currentStep := CurrentStep{heaterOn: true}
	heaterOn := currentStep.heaterHysteresis(tempComp)
	if !heaterOn || !currentStep.heaterOn {
		t.Error("Heater should be on")
	}
}

func TestCheckHeaterOnWhileOk(t *testing.T) {
	tempComp := OK
	currentStep := CurrentStep{heaterOn: false}
	heaterOn := currentStep.heaterHysteresis(tempComp)
	if heaterOn || currentStep.heaterOn {
		t.Error("Heater should be off")
	}
}
