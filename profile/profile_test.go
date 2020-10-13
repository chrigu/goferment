package profile

import (
	"math"
	"testing"
	"time"
)

// step & currentstep tests

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

func TestStepActivation(t *testing.T) {
	currentStep := CurrentStep{active: false}

	currentStep.activateStep()
	if !currentStep.active || currentStep.startTime.IsZero() {
		t.Error("CurrentStep should be activated")
	}
}

func TestCurrentStepTimeLeft(t *testing.T) {
	minuteDelta := 32
	duration := 69
	pastDate := time.Now().Add(-time.Duration(minuteDelta) * time.Minute)
	currentStep := CurrentStep{startTime: pastDate, ProfileStep: &ProfileStep{Duration: duration}}

	if math.Round(currentStep.stepTimeLeft()) != float64(duration-minuteDelta) {
		t.Error("Could not calculate time left")
	}

	if currentStep.hasEnded() {
		t.Error("Step has not yet ended")
	}
}

func TestCurrentStepHasEnded(t *testing.T) {
	minuteDelta := 11
	duration := 10
	pastDate := time.Now().Add(-time.Duration(minuteDelta) * time.Minute)
	currentStep := CurrentStep{startTime: pastDate, ProfileStep: &ProfileStep{Duration: duration}}

	if !currentStep.hasEnded() {
		t.Error("Step has ended")
	}
}

// json test
func TestProfileFromJsonFile(t *testing.T) {
	profile := ReadProfileFromFile("test-profile.json")
	if profile.Steps[0].Name != "Step 1" || profile.Steps[1].Name != "Step 2" || profile.Steps[2].Name != "Step 3" {
		t.Error("Steps were not imported correctly")
	}
}
