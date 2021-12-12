package basins

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
