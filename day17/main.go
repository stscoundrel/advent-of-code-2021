package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day17/probes"
)

func CalculateHighestShot(targetGrid probes.Grid) int {
	result, _ := probes.BruteforceHighestShot(targetGrid)

	return result
}

func CalculateDistinctHits(targetGrid probes.Grid) int {
	_, result := probes.BruteforceHighestShot(targetGrid)

	return result
}
