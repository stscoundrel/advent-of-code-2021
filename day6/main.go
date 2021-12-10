package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day6/fishcounter"
	"github.com/stscoundrel/advent-of-code-2021/day6/reader"
)

func CountFishes() int {
	numbers := reader.ReadIntsFromFile("./resources/input.txt")
	fishes := fishcounter.NewFishGroup(numbers)
	count := fishcounter.GetFishCount(fishes, 80)

	return count
}
