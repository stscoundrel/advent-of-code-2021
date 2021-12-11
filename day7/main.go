package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day7/reader"
	"github.com/stscoundrel/advent-of-code-2021/day7/submarines"
)

func GetSmallestFuelConsumption(linear bool) int {
	numbers := reader.ReadIntsFromFile("./resources/input.txt")
	crabSubmarines := submarines.NewSubmarineGroup(numbers)
	consumption := submarines.CalculateAlignConsumption(crabSubmarines, linear)

	return consumption
}
