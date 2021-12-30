package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day17/probes"
)

func CalculateHighestShot(targetGrid probes.Grid) int {
	return probes.BruteforceHighestShot(targetGrid)
}
