package main

import (
	"fmt"

	"github.com/stscoundrel/advent-of-code-2021/day5/coordinates"
	"github.com/stscoundrel/advent-of-code-2021/day5/reader"
)

func GetOverLappingWithTestData() int {
	rawLines := reader.ReadLinesFromFile("./reader/test_input.txt")
	lines := coordinates.ParseCoordinates(rawLines)
	overlapping := coordinates.GetOverLappingPoints(lines)

	return overlapping
}

func GetOverLapping() int {
	rawLines := reader.ReadLinesFromFile("./resources/vents.txt")
	lines := coordinates.ParseCoordinates(rawLines)
	overlapping := coordinates.GetOverLappingPoints(lines)

	return overlapping
}

func main() {
	fmt.Println(GetOverLappingWithTestData())
	fmt.Println(GetOverLapping())
}
