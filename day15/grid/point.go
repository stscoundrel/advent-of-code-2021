package grid

type Point struct {
	x     int
	y     int
	value int
}

func (point *Point) hasNeighbour(newX int, newY int, grid [][]int) bool {
	return pointExistsInMap(newX, newY, grid)
}

func (point *Point) hasNeighbourUp(grid [][]int) bool {
	newX := point.x - 1
	newY := point.y

	return point.hasNeighbour(newX, newY, grid)
}

func (point *Point) hasNeighbourDown(grid [][]int) bool {
	newX := point.x + 1
	newY := point.y

	return point.hasNeighbour(newX, newY, grid)
}

func (point *Point) hasNeighbourLeft(grid [][]int) bool {
	newX := point.x
	newY := point.y - 1

	return point.hasNeighbour(newX, newY, grid)
}

func (point *Point) hasNeighbourRight(grid [][]int) bool {
	newX := point.x
	newY := point.y + 1

	return point.hasNeighbour(newX, newY, grid)
}

func NewPoint(x int, y int, value int) Point {
	return Point{x: x, y: y, value: value}
}
