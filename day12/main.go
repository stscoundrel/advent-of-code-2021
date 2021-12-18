package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day12/paths"
	"github.com/stscoundrel/advent-of-code-2021/day12/reader"
)

func CountPaths(fileName string, steps int) int {
	pathInput := reader.ReadPathsFromFile(fileName)
	count := paths.CountPaths(pathInput)

	return count
}
