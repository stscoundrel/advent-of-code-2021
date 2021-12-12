package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day9/basins"
	"github.com/stscoundrel/advent-of-code-2021/day9/heights"
	"github.com/stscoundrel/advent-of-code-2021/day9/reader"
)

func CalculateRisks(fileName string) int {
	heightMap := reader.ReadHeighMapFromFile(fileName)
	lowpoints := heights.GetLowPoints(heightMap)
	risk := heights.CalculateRiskLevel(lowpoints)

	return risk
}

func CalculateThreeLargest(fileName string) int {
	heightMap := reader.ReadHeighMapFromFile(fileName)
	filledBasins := basins.GetBasins(heightMap)
	multiplied := basins.MultiplyThreeLargest(filledBasins)

	return multiplied
}
