package boardgame

import (
	"testing"
)

func TestScoresBingoBoard(t *testing.T) {
	boards := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}

	noWinNumbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21,
	}

	winNumbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24,
	}

	loseResult := CheckBoard(boards, noWinNumbers)
	winResult := CheckBoard(boards, winNumbers)

	if loseResult != false {
		t.Error("Did not return expected content. Received", loseResult, "expected ", false)
	}

	if winResult != true {
		t.Error("Did not return expected content. Received", winResult, "expected ", true)
	}

}

func TestScoresBingoBoards(t *testing.T) {
	boards := [][][]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}

	numbers := []int{
		7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
	}

	result := Play(boards, numbers)
	expected := 4512

	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected ", expected)
	}

}
