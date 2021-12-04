package diagnostics

import (
	"testing"
)

func TestGetsPowerConsumption(t *testing.T) {

	result := CalculatePowerConsumption(22, 9)
	var expected int64 = 198

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
