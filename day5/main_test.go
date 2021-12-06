package main

import (
	"testing"
)

func TestGetOverlappingWithTestData(t *testing.T) {
	result := GetOverLappingWithTestData()
	var expected int = 5

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetOverlappingDiagonalWithTestData(t *testing.T) {
	result := GetOverLappingDiagonalWithTestData()
	var expected int = 12

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetOverlapping(t *testing.T) {
	result := GetOverLapping()
	var expected int = 5145

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetOverlappingDiagonal(t *testing.T) {
	result := GetOverLappingDiagonal()
	var expected int = 16518

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
