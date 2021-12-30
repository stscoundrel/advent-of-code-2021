package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day16/reader"
	"github.com/stscoundrel/advent-of-code-2021/day16/transmissions"
)

func SumVersionNumbers(fileName string) int {
	hexTransmission := reader.ReadGTransmissionFromFile(fileName)
	versionSum := transmissions.SumVersionNumbers(hexTransmission)

	return versionSum
}
