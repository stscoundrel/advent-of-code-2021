package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day14/polymer"
	"github.com/stscoundrel/advent-of-code-2021/day14/reader"
)

func GetMostAndLeastCommonDifference(fileName string, iterations int) int {
	template, rawRules := reader.ReadInstructions(fileName)
	rules := polymer.ParseRules(rawRules)
	most := 0
	least := 0

	// Use linear bruteforce for part 1
	if iterations == 10 {
		most, least = polymer.GetMostAndLeastCommonLinear(template, rules, iterations)
		return most - least
	} else {
		most, least = polymer.GetMostAndLeastCommon(template, rules, iterations)
		return most - least
	}
}
