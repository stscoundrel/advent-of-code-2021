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

func BruteforceHighestShot(grid Grid) (int, int) {
	xLimit := 275
	yLimit := 275
	succesfulProbes := []Probe{}

	for x := -xLimit; x < xLimit; x += 1 {
		for y := -yLimit; y < yLimit; y += 1 {
			probeIsInTheAir := true
			probe := newProbe(x, y)

			for probeIsInTheAir {
				probe.move()

				if probe.isPastGrid(grid) {
					probeIsInTheAir = false
					continue
				}

				if probe.hitGrid(grid) {
					succesfulProbes = append(succesfulProbes, probe)
					probeIsInTheAir = false
				}
			}
		}
	}

	return getHighestProbe(succesfulProbes), len(succesfulProbes)
}
