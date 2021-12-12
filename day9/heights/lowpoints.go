package heights

func IsLowPoint(i int, j int, heights [][]int) bool {
	row := heights[i]
	current := row[j]
	isFirstRow := i == 0
	isLastRow := i == len(heights)-1
	isFirstOfRow := j == 0
	isLastOfRow := j == len(row)-1

	// First row should only one lower.
	if isFirstRow {
		nextRow := heights[i+1]

		if current >= nextRow[j] {
			return false
		}
	}

	// Last row should check upper
	if isLastRow {
		previousRow := heights[i-1]

		if current >= previousRow[j] {
			return false
		}
	}

	// All expect first of row compared to previouus one
	if !isFirstOfRow {
		if current >= row[j-1] {
			return false
		}
	}

	// All except last of row compared to the next one
	if !isLastOfRow {
		if current >= row[j+1] {
			return false
		}
	}

	// Middle rows should compare to upper & lower
	if !isFirstRow && !isLastRow {
		previousRow := heights[i-1]
		nextRow := heights[i+1]

		if current >= previousRow[j] || current >= nextRow[j] {
			return false
		}
	}

	return true
}

func GetLowPoints(heights [][]int) []int {
	lowPoints := []int{}

	for i := 0; i < len(heights); i += 1 {
		row := heights[i]
		for j := 0; j < len(row); j += 1 {
			if IsLowPoint(i, j, heights) {
				lowPoints = append(lowPoints, heights[i][j])
			}
		}
	}

	return lowPoints
}
