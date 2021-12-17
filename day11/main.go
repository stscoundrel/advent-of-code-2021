package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day11/octopuses"
	"github.com/stscoundrel/advent-of-code-2021/day11/reader"
)

func CalculateFlashes(fileName string, steps int) int {
	grid := reader.ReadGridFromFile(fileName)
	octopusGroup := octopuses.GridToOctopuses(grid)
	flashes := octopuses.PassTime(steps, octopusGroup)

	return flashes
}
