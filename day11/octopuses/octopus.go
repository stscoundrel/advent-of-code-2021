package octopuses

type Octopus struct {
	x         int
	y         int
	energy    int
	lastFlash int
}

func pointHasOctopus(x int, y int, octopuses [][]Octopus) bool {
	validX := true
	validY := true

	if x < 0 || len(octopuses)-1 < x {
		return false
	}

	if y < 0 || len(octopuses[x])-1 < y {
		return false
	}

	return validX && validY
}

func (octopus *Octopus) flashUpLeft(step int, octopuses [][]Octopus) {
	newX := octopus.x + 1
	newY := octopus.y - 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashUpRight(step int, octopuses [][]Octopus) {
	newX := octopus.x + 1
	newY := octopus.y + 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashUp(step int, octopuses [][]Octopus) {
	newX := octopus.x + 1
	newY := octopus.y

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashDownLeft(step int, octopuses [][]Octopus) {
	newX := octopus.x - 1
	newY := octopus.y - 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashDownRight(step int, octopuses [][]Octopus) {
	newX := octopus.x - 1
	newY := octopus.y + 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashDown(step int, octopuses [][]Octopus) {
	newX := octopus.x - 1
	newY := octopus.y

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashLeft(step int, octopuses [][]Octopus) {
	newX := octopus.x
	newY := octopus.y - 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) flashRight(step int, octopuses [][]Octopus) {
	newX := octopus.x
	newY := octopus.y + 1

	if pointHasOctopus(newX, newY, octopuses) {
		octopuses[newX][newY].increaseEnergy(step, octopuses)
	}
}

func (octopus *Octopus) increaseEnergy(step int, octopuses [][]Octopus) {
	if octopus.lastFlash != step {
		if octopus.energy+1 > 9 {
			octopus.energy = 0
			octopus.lastFlash = step
			octopus.flash(step, octopuses)

		} else {
			octopus.energy += 1
		}
	}
}

func (octopus *Octopus) flash(step int, octopuses [][]Octopus) {
	octopus.flashUpLeft(step, octopuses)
	octopus.flashUpRight(step, octopuses)
	octopus.flashUp(step, octopuses)
	octopus.flashDown(step, octopuses)
	octopus.flashDownLeft(step, octopuses)
	octopus.flashDownRight(step, octopuses)
	octopus.flashLeft(step, octopuses)
	octopus.flashRight(step, octopuses)
}
