package octopuses

func getFlashes(step int, octopuses [][]Octopus) int {
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

func PassTime(steps int, octopuses [][]Octopus) int {
	flashes := 0

	for step := 0; step <= steps; step += 1 {
		for x, row := range octopuses {
			for y, _ := range row {
				octopuses[x][y].increaseEnergy(step, octopuses)
			}
		}

		flashes += getFlashes(step, octopuses)

		// for x, row := range octopuses {
		// 	for y, _ := range row {
		// 		fmt.Print(octopuses[x][y].energy)
		// 	}
		// 	fmt.Println("")
		// }
		// fmt.Println("#############################")
	}

	return flashes
}
