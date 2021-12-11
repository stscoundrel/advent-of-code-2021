package submarines

func getSmallestConsumption(consumptions []int) int {
	smallest := consumptions[0]

	for _, consumption := range consumptions {
		if consumption < smallest {
			smallest = consumption
		}
	}

	return smallest
}

func getConsumptionTo(position int, submarines []Submarine, linear bool) int {
	consumption := 0

	for _, submarine := range submarines {
		if linear {
			consumption += submarine.getLinearFuelConsumption(position)
		} else {
			consumption += submarine.getIncreasingFuelConsumption(position)
		}
	}

	return consumption
}

func CalculateAlignConsumption(submarines []Submarine, linear bool) int {
	consumptions := []int{}
	start := submarines[0].position
	end := submarines[len(submarines)-1].position

	for position := start; position < end; position += 1 {
		consumptions = append(consumptions, getConsumptionTo(position, submarines, linear))
	}

	return getSmallestConsumption(consumptions)
}
