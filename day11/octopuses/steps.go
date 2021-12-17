package octopuses

func getFlashes(octopuses [][]Octopus) int {
	flashes := 0

	for x, row := range octopuses {
		for y, _ := range row {
			if octopuses[x][y].energy == 0 {
				flashes += 1
			}
		}
	}

	return flashes
}

func isSimulatenousFlash(octopuses [][]Octopus) bool {
	return getFlashes(octopuses) == 100
}

func PassTime(steps int, octopuses [][]Octopus) int {
	flashes := 0

	for step := 0; step <= steps; step += 1 {
		for x, row := range octopuses {
			for y, _ := range row {
				octopuses[x][y].increaseEnergy(step, octopuses)
			}
		}

		flashes += getFlashes(octopuses)
	}

	return flashes
}

func GetSimulatenousFlash(octopuses [][]Octopus) int {
	step := 1
	simultaneousFlash := 0

	for simultaneousFlash == 0 {
		for x, row := range octopuses {
			for y, _ := range row {
				octopuses[x][y].increaseEnergy(step, octopuses)
			}
		}

		if isSimulatenousFlash(octopuses) {
			simultaneousFlash = step
		}

		step += 1
	}

	return simultaneousFlash
}
