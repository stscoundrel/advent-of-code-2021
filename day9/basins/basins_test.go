package basins

import (
	"testing"
)

func TestGetsInitialBasins(t *testing.T) {
	heights := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	result := GetInitialBasins(heights)

	if len(result) != 4 {
		t.Error("Result did not match length of expected. Got", len(result), "expected ", 4)
	}
}

func TestFillsOutABasin(t *testing.T) {
	heights := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	point := NewPoint(0, 1, 1)
	points := []Point{point}
	basin := Basin{points: points}

	basin = FillABasin(basin, heights)

	if len(basin.points) != 3 {
		t.Error("Result did not match length of expected. Got", len(basin.points), "expected ", 3)
	}
}

func TestFillsOutBasins(t *testing.T) {
	heights := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	initialBasins := GetInitialBasins(heights)
	basins := FillBasins(initialBasins, heights)

	expectedSizes := []int{
		3, 9, 14, 9,
	}

	for index, basin := range basins {
		if len(basin.points) != expectedSizes[index] {
			t.Error("Result did not match length of expected. Got", len(basin.points), "expected ", expectedSizes[index])
		}
	}
}

func TestMultipliesThreeLargest(t *testing.T) {
	heights := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	initialBasins := GetInitialBasins(heights)
	basins := FillBasins(initialBasins, heights)
	result := MultiplyThreeLargest(basins)
	expected := 1134

	if result != expected {
		t.Error("Result did not match length of expected. Got", result, "expected ", expected)
	}

}
