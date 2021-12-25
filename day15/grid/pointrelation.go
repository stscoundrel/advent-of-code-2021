package grid

import "strconv"

type PointRelation struct {
	Origin      string
	Destination string
	Value       int
}

func NewNeighbour(oldPoint Point, x int, y int, value int) PointRelation {
	return PointRelation{
		Origin:      strconv.Itoa(oldPoint.x) + string('-') + strconv.Itoa(oldPoint.y),
		Destination: strconv.Itoa(x) + string('-') + strconv.Itoa(y),
		Value:       value,
	}
}

func NewNeighbourUp(point Point, heighMap [][]int) PointRelation {
	newX := point.x - 1
	newY := point.y

	return NewNeighbour(point, newX, newY, heighMap[newX][newY])
}

func NewNeighbourDown(point Point, heighMap [][]int) PointRelation {
	newX := point.x + 1
	newY := point.y

	return NewNeighbour(point, newX, newY, heighMap[newX][newY])
}

func NewNeighbourLeft(point Point, heighMap [][]int) PointRelation {
	newX := point.x
	newY := point.y - 1

	return NewNeighbour(point, newX, newY, heighMap[newX][newY])
}

func NewNeighbourRight(point Point, heighMap [][]int) PointRelation {
	newX := point.x
	newY := point.y + 1

	return NewNeighbour(point, newX, newY, heighMap[newX][newY])
}
