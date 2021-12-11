package submarines

import (
	"testing"
)

func TestCalculatesFuelConsumptionOnAlign(t *testing.T) {
	initialPositions := []int{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	submarines := NewSubmarineGroup((initialPositions))

	result := CalculateAlignConsumption(submarines, true)

	expected := 37

	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected ", expected)
	}
}

func TestCalculatesIncreasingFuelConsumptionOnAlign(t *testing.T) {
	initialPositions := []int{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	submarines := NewSubmarineGroup((initialPositions))

	result := CalculateAlignConsumption(submarines, false)

	expected := 168

	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected ", expected)
	}
}
