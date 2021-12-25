package grid

func pointExistsInMap(x int, y int, mapping [][]int) bool {
	validX := true
	validY := true

	if x < 0 || len(mapping)-1 < x {
		return false
	}

	if y < 0 || len(mapping[x])-1 < y {
		return false
	}

	return validX && validY
}

func GetPointNeighbours(point Point, grid [][]int) []PointRelation {
	neighbours := []PointRelation{}

	if point.hasNeighbourUp(grid) {
		neighbours = append(neighbours, NewNeighbourUp(point, grid))
	}

	if point.hasNeighbourDown(grid) {
		neighbours = append(neighbours, NewNeighbourDown(point, grid))
	}

	if point.hasNeighbourLeft(grid) {
		neighbours = append(neighbours, NewNeighbourLeft(point, grid))
	}

	if point.hasNeighbourRight(grid) {
		neighbours = append(neighbours, NewNeighbourRight(point, grid))
	}

	return neighbours
}

func GridToEdges(grid [][]int) {

}
