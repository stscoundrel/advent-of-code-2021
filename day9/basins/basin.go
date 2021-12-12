package basins

type Basin struct {
	points []Point
}

func (basin *Basin) isFull() bool {
	hasNonExpandedPoints := true

	for _, point := range basin.points {
		if !point.hasExpanded {
			hasNonExpandedPoints = false
		}
	}

	return hasNonExpandedPoints
}

func (basin *Basin) hasPoint(x int, y int) bool {
	for _, point := range basin.points {
		if point.x == x && point.y == y {
			return true
		}
	}

	return false
}

func (basin *Basin) getSize() int {
	return len(basin.points)
}
