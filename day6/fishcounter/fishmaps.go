package fishcounter

func getInitialFishMap(fishes []int) [9]int {
	var fishMap [9]int

	for _, number := range fishes {
		fishMap[number] += 1
	}

	return fishMap
}

func sumFishMap(fishes [9]int) int {
	sum := 0

	for _, amount := range fishes {
		sum += amount
	}

	return sum
}
