package measurements

func CountValueIncreases(values []int) (increases int) {
	increases = 0

	for index := range values {
		if index == 0 {
			continue
		}

		if values[index] > values[index-1] {
			increases += 1
		}
	}

	return
}
