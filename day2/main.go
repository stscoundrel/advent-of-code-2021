package main

import (
	"fmt"

	"github.com/stscoundrel/advent-of-code-2021/day2/locator"
	"github.com/stscoundrel/advent-of-code-2021/day2/reader"
)

func LocateShip() int {
	steps := reader.ReadStringsFromFile("./resources/steps.txt")
	location := locator.Locate(steps)

	return location
}

func main() {
	fmt.Println(LocateShip())
}
