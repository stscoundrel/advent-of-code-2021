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
