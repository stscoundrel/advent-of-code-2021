package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day6/fishcounter"
	"github.com/stscoundrel/advent-of-code-2021/day6/reader"
)

func CountFishes(days int) int {
	numbers := reader.ReadIntsFromFile("./resources/input.txt")
	fishes := fishcounter.NewFishGroup(numbers)
	count := fishcounter.GetFishCount(fishes, days)

	return count
}

func CountFishesInGroups(days int) int {
	numbers := reader.ReadIntsFromFile("./resources/input.txt")
	count := fishcounter.GetFishGroupCount(numbers, days)

	return count
}
