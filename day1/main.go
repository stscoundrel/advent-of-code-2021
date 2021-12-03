package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day1/measurements"
	"github.com/stscoundrel/advent-of-code-2021/day1/reader"
)

func GetIncreasesInDepth() int {
	depthValues := reader.ReadIntsFromFile("./resources/depths.txt")
	increases := measurements.CountValueIncreases(depthValues)

	return increases
}

func GetIncreasesInDepthSliding() int {
	depthValues := reader.ReadIntsFromFile("./resources/depths.txt")
	increases := measurements.CountValueIncreasesSliding(depthValues)

	return increases
}
