package polymer

func getMostAndLeastCommonInMap(frequencyMap map[string]int) (int, int) {
	most := 0
	least := 0

	for _, value := range frequencyMap {
		if least == 0 && most == 0 {
			least = value
			most = value
		}

		if value < least {
			least = value
		}

		if value > most {
			most = value
		}
	}

	return most, least
}

func getMostAndLeastCommonInPairMap(pairsMap map[string]int) (int, int) {
	singlesMap := map[string]int{}

	for pair, count := range pairsMap {
		singlesMap[string(pair[0])] += count
	}

	return getMostAndLeastCommonInMap(singlesMap)
}
