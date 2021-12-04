package main

import (
	"testing"
)

func TestGetPowerConsumption(t *testing.T) {
	result := ReadPowerConsumption()
	var expected int64 = 1071734

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
