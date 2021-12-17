package octopuses

func newOctopus(x int, y int, energy int) Octopus {
	return Octopus{x: x, y: y, energy: energy}
}

func GridToOctopuses(grid [][]int) [][]Octopus {
	octopuses := [][]Octopus{}

	for x, row := range grid {
		octopusRow := []Octopus{}

		for y, number := range row {
			octopusRow = append(octopusRow, newOctopus(x, y, number))
		}

		octopuses = append(octopuses, octopusRow)
	}

	return octopuses
}
