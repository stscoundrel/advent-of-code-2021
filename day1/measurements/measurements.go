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

func CountValueIncreasesSliding(values []int) int {
	increases := 0

	for index := 0; index < len(values)-3; index += 1 {
		sum := values[index] + values[index+1] + values[index+2]
		nextSum := values[index+1] + values[index+2] + values[index+3]

		if nextSum > sum {
			increases += 1
		}
	}

	return increases
}
