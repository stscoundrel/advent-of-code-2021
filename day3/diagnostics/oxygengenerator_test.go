package diagnostics

import (
	"testing"
)

func TestGetsOxygenRateBinary(t *testing.T) {
	steps := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := GetOxygenRateBinary(steps)
	expected := "10111"

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}

func TestGetsOxygenRateDecimal(t *testing.T) {
	steps := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := GetOxygenRateDecimal(steps)
	var expected int64 = 23

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
