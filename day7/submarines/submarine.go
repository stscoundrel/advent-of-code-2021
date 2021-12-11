package submarines

import "sort"

type Submarine struct {
	position int
}

func (submarine Submarine) getLinearFuelConsumption(destination int) int {
	if submarine.position > destination {
		return submarine.position - destination
	}

	return destination - submarine.position
}

func (submarine Submarine) getIncreasingFuelConsumption(destination int) int {
	consumption := 0
	previousConsumption := 0
	distance := submarine.getLinearFuelConsumption(destination)

	for i := 0; i < distance; i += 1 {
		consumption += previousConsumption + 1
		previousConsumption += 1
	}

	return consumption
}

func NewSubmarineGroup(positions []int) []Submarine {
	sort.Ints(positions)
	submarines := []Submarine{}

	for _, number := range positions {
		submarines = append(submarines, Submarine{number})
	}

	return submarines
}
