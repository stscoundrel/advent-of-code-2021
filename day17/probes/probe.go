package probes

type Probe struct {
	x         int
	y         int
	xVelocity int
	yVelocity int
	highestY  int
}

func (packet *Probe) move() {
	packet.x += packet.xVelocity
	packet.y += packet.yVelocity

	if packet.y > packet.highestY {
		packet.highestY = packet.y
	}

	if packet.xVelocity > 0 {
		packet.xVelocity -= 1
	} else {
		packet.xVelocity += 1
	}

	packet.yVelocity = packet.yVelocity - 1
}

func (packet *Probe) hitGrid(grid Grid) bool {
	if packet.x >= grid.XStart && packet.x <= grid.XEnd && packet.y <= grid.YStart && packet.y >= grid.YEnd {
		return true
	}

	return false
}

func (packet *Probe) isPastGrid(grid Grid) bool {
	if packet.x > grid.XEnd || packet.y < grid.YEnd {
		return true
	}

	return false
}

func newProbe(x int, y int) Probe {
	return Probe{
		x:         0,
		y:         0,
		xVelocity: x,
		yVelocity: y,
		highestY:  0,
	}
}
