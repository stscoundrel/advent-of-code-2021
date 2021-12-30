package main

import (
	"testing"

	"github.com/stscoundrel/advent-of-code-2021/day17/probes"
)

func TestGetsHighestInTestGrid(t *testing.T) {
	grid := probes.Grid{XStart: 20, XEnd: 30, YStart: -5, YEnd: -10}
	result := CalculateHighestShot(grid)
	expected := 45

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsHighestInGrid(t *testing.T) {
	grid := probes.Grid{XStart: 207, XEnd: 263, YStart: -63, YEnd: -115}
	result := CalculateHighestShot(grid)
	expected := 6555

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
