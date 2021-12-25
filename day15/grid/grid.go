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

func Enlarge(grid [][]int) [][]int {
	multiplier := 5

	width := len(grid[0])
	height := len(grid)
	newWidth := multiplier * width
	newHeight := multiplier * height

	largerGrid := [][]int{}

	for i := 0; i < newHeight; i += 1 {
		largerGrid = append(largerGrid, make([]int, newWidth))
	}

	for xM := 0; xM < 5; xM++ {
		for yM := 0; yM < 5; yM++ {
			for xT := 0; xT < width; xT++ {
				for yT := 0; yT < height; yT++ {
					x := xM*width + xT
					y := yM*height + yT
					value := grid[xT][yT] + xM + yM

					largerGrid[x][y] = ((value - 1) % 9) + 1
				}
			}
		}
	}
	return largerGrid
}
