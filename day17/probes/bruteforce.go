package probes

func getHighestProbe(probes []Probe) int {
	highest := probes[0].highestY

	for _, probe := range probes {
		if probe.highestY > highest {
			highest = probe.highestY
		}
	}

	return highest
}

func BruteforceHighestShot(grid Grid) int {
	xLimit := 150
	yLimit := 150
	succesfulProbes := []Probe{}

	for x := 1; x < xLimit; x += 1 {
		for y := 1; y < yLimit; y += 1 {
			probeIsInTheAir := true
			probe := newProbe(x, y)

			for probeIsInTheAir {
				probe.move()

				if probe.hitGrid(grid) {
					succesfulProbes = append(succesfulProbes, probe)
					probeIsInTheAir = false
				}

				if probe.isPastGrid(grid) {
					probeIsInTheAir = false
				}
			}
		}
	}

	return getHighestProbe(succesfulProbes)
}
