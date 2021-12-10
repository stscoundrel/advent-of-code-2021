package fishcounter

func GetFishCount(initialFishes []Fish, days int) int {
	fishes := initialFishes

	for day := 0; day < days; day += 1 {
		for i, fish := range fishes {
			fishes[i] = fish.passADay()
		}

		for _, fish := range fishes {
			if fish.shouldSpawn() {
				fishes = append(fishes, newFish(8))
			}
		}
	}

	return len(fishes)
}

func GetFishGroupCount(initialFishes []int, days int) int {
	fishMap := getInitialFishMap(initialFishes)

	for day := 0; day < days; day += 1 {
		var newFishMap [9]int

		for group, fishAmount := range fishMap {
			if group == 0 {
				// Change zeros to sixes
				newFishMap[6] += fishAmount

				// Append corresponding amount of new fishes.
				newFishMap[8] += fishAmount
			} else {
				// Age fishes to the previous group.
				newFishMap[group-1] += fishAmount
			}
		}

		fishMap = newFishMap
	}

	return sumFishMap(fishMap)
}
