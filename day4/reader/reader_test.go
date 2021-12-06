package reader

import (
	"testing"
)

func TestReadsBingoNumbers(t *testing.T) {
	result := ReadBingoNumbersFromFile("./test_input.txt")
	expected := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}

func TestReadsBingoBoards(t *testing.T) {
	result := ReadBingoBoardsFromFile("./test_input.txt")
	expected1 := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	expected2 := [][]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}

	expected3 := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}

	for rowIndex, row := range expected1 {
		for columnIndex, _ := range row {
			if result[0][rowIndex][columnIndex] != expected1[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", result[0][rowIndex][columnIndex], "expected ", expected1[rowIndex][columnIndex])
			}
		}
	}

	for rowIndex, row := range expected2 {
		for columnIndex, _ := range row {
			if result[1][rowIndex][columnIndex] != expected2[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", result[1][rowIndex][columnIndex], "expected ", expected2[rowIndex][columnIndex])
			}
		}
	}

	for rowIndex, row := range expected3 {
		for columnIndex, _ := range row {
			if result[2][rowIndex][columnIndex] != expected3[rowIndex][columnIndex] {
				t.Error("Did not return expected content. Received", result[2][rowIndex][columnIndex], "expected ", expected3[rowIndex][columnIndex])
			}
		}
	}
}
