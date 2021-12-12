package heights

import (
	"testing"
)

func TestGetsLowPoints(t *testing.T) {
	heights := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	expected := []int{
		1, 0, 5, 5,
	}

	result := GetLowPoints(heights)

	if len(result) != len(expected) {
		t.Error("Result did not match length of expected. Got", len(result), "expected ", len(expected))
	}

	for index, value := range result {
		if value != expected[index] {
			t.Error("Did not return expected content. Received", value, "expected ", expected[index])
		}
	}

}
