package main

import (
	"testing"
)

func TestGetsFuelConsumption(t *testing.T) {
	result := GetSmallestFuelConsumption(true)
	var expected int = 337488

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsIncreasingFuelConsumption(t *testing.T) {
	result := GetSmallestFuelConsumption(false)
	var expected int = 89647695

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
