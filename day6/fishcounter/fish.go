package fishcounter

type Fish struct {
	current  int
	previous int
}

func (fish Fish) passADay() Fish {
	fish.previous = fish.current

	if fish.current-1 >= 0 {
		fish.current -= 1
	} else {
		fish.current = 6
	}

	return fish
}

func (fish *Fish) shouldSpawn() bool {
	return fish.current == 6 && fish.previous == 0
}

func newFish(day int) Fish {
	return Fish{current: day, previous: day}
}

func NewFishGroup(days []int) []Fish {
	fishes := []Fish{}

	for _, number := range days {
		fishes = append(fishes, newFish(number))
	}

	return fishes
}
