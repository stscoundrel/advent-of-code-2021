package basins

func NewPoint(x int, y int, value int) Point {
	return Point{x: x, y: y, value: value, hasExpanded: false}
}

func NewPointUp(point Point, heighMap [][]int) Point {
	newX := point.x - 1
	newY := point.y

	return NewPoint(newX, newY, heighMap[newX][newY])
}

func NewPointDown(point Point, heighMap [][]int) Point {
	newX := point.x + 1
	newY := point.y

	return NewPoint(newX, newY, heighMap[newX][newY])
}

func NewPointLeft(point Point, heighMap [][]int) Point {
	newX := point.x
	newY := point.y - 1

	return NewPoint(newX, newY, heighMap[newX][newY])
}

func NewPointRight(point Point, heighMap [][]int) Point {
	newX := point.x
	newY := point.y + 1

	return NewPoint(newX, newY, heighMap[newX][newY])
}
