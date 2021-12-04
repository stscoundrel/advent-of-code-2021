package main

import (
	"fmt"

	"github.com/stscoundrel/advent-of-code-2021/day3/diagnostics"
	"github.com/stscoundrel/advent-of-code-2021/day3/reader"
)

func ReadPowerConsumption() int64 {
	values := reader.ReadStringsFromFile("./resources/diagnostics.txt")
	powerConsumption := diagnostics.GetPowerConsumptionFromDiagnostics(values)

	return powerConsumption
}

func main() {
	fmt.Println(ReadPowerConsumption())
}
