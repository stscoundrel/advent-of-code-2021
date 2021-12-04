package diagnostics

import (
	"testing"
)

func TestGetsGammaRateBinary(t *testing.T) {
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

	result := GetGammaRateBinary(steps)
	expected := "10110"

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}

func TestGetsGammaRateDecimal(t *testing.T) {
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

	result := GetGammaRateDecimal(steps)
	var expected int64 = 22

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
