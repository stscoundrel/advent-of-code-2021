package reader

import (
	"fmt"
	"testing"
)

func TestReadsSignalInputFromFile(t *testing.T) {
	result := ReadGridFromFile("./test_input.txt")
	expected := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	fmt.Println(result)

	for rowIndex, row := range result {
		for columnIndex, _ := range row {
			if result[rowIndex][columnIndex] != expected[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", result[rowIndex][columnIndex], "expected ", expected[rowIndex][columnIndex])
			}
		}
	}
}
