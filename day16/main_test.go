package main

import "testing"

func TestSumsVersionNumbers(t *testing.T) {
	result := SumVersionNumbers("./resources/input.txt")
	var expected int = 821

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsPackageValue(t *testing.T) {
	result := GetTransmissionValue("./resources/input.txt")
	var expected int = 2056021084691

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
