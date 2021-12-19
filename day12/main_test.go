package main

import (
	"testing"
)

func TestCountsPathsInTestInput(t *testing.T) {
	result := CountPaths("./reader/test_input.txt")
	var expected int = 10

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsPathsInTestInput2(t *testing.T) {
	result := CountPaths("./reader/test_input2.txt")
	var expected int = 19

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsPathsInTestInput3(t *testing.T) {
	result := CountPaths("./reader/test_input3.txt")
	var expected int = 226

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsPathsInInput(t *testing.T) {
	result := CountPaths("./resources/input.txt")
	var expected int = 4773

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsDualVisitPathsInTestInput(t *testing.T) {
	result := CountDualVisitPaths("./reader/test_input.txt")
	var expected int = 36

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsDualVisitPathsInTestInput2(t *testing.T) {
	result := CountDualVisitPaths("./reader/test_input2.txt")
	var expected int = 103

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsDualVisitPathsInTestInput3(t *testing.T) {
	result := CountDualVisitPaths("./reader/test_input3.txt")
	var expected int = 3509

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsDualVisitPathsInInput(t *testing.T) {
	result := CountDualVisitPaths("./resources/input.txt")
	var expected int = 116985

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
