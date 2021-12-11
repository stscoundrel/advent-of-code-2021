package heights

func CalculateRiskLevel(lowpoints []int) int {
	sum := 0

	for _, value := range lowpoints {
		sum += value + 1
	}

	return sum
}
