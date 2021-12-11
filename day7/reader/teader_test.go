package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadIntsFromFile("./test_input.txt")
	expected := []int{
		16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
