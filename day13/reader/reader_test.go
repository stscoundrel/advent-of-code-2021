package reader

import (
	"testing"
)

func TestReadsHeightMaps(t *testing.T) {
	points, folds := ReadFoldInstructions("./test_input.txt")
	expectedPoints := [][]int{
		{6, 10},
		{0, 14},
		{9, 10},
		{0, 3},
		{10, 4},
		{4, 11},
		{6, 0},
		{6, 12},
		{4, 1},
		{0, 13},
		{10, 12},
		{3, 4},
		{3, 0},
		{8, 4},
		{1, 10},
		{2, 14},
		{8, 10},
		{9, 0},
	}

	expectedFolds := []string{
		"fold along y=7",
		"fold along x=5",
	}

	for rowIndex, row := range points {
		for columnIndex, column := range row {
			if column != expectedPoints[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", column, "expected ", expectedPoints[rowIndex][columnIndex])
			}
		}
	}

	for index, fold := range folds {
		if fold != expectedFolds[index] {
			t.Error("Did not return expected content. Received", fold, "expected ", expectedFolds[index])
		}
	}

}
