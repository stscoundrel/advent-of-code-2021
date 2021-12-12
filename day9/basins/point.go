package basins

type Point struct {
	x           int
	y           int
	value       int
	hasExpanded bool
}

func (point *Point) canExpand(newX int, newY int, heightMap [][]int, basin Basin) bool {
	if !pointExistsInMap(newX, newY, heightMap) {
		return false
	}

	if basin.hasPoint(newX, newY) {
		return false
	}

	if heightMap[newX][newY] == 9 {
		return false
	}

	return true
}

func (point *Point) canExpandUp(heightMap [][]int, basin Basin) bool {
	newX := point.x - 1
	newY := point.y

	return point.canExpand(newX, newY, heightMap, basin)
}

func (point *Point) canExpandDown(heightMap [][]int, basin Basin) bool {
	newX := point.x + 1
	newY := point.y

	return point.canExpand(newX, newY, heightMap, basin)
}

func (point *Point) canExpandLeft(heightMap [][]int, basin Basin) bool {
	newX := point.x
	newY := point.y - 1

	return point.canExpand(newX, newY, heightMap, basin)
}

func (point *Point) canExpandRight(heightMap [][]int, basin Basin) bool {
	newX := point.x
	newY := point.y + 1

	return point.canExpand(newX, newY, heightMap, basin)
}

func (point *Point) shouldExpand() bool {
	return !point.hasExpanded
}
