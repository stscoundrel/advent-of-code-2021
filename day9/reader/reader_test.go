package reader

import (
	"testing"
)

func TestReadsHeightMaps(t *testing.T) {
	result := ReadHeighMapFromFile("./test_input.txt")
	expected := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	for rowIndex, row := range result {
		for columnIndex, column := range row {
			if column != expected[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", column, "expected ", expected[rowIndex][columnIndex])
			}
		}
	}

}
