package locator

import (
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	length    int
}

func stepsToMovements(steps []string) []Movement {
	var movements []Movement

	for _, step := range steps {
		parts := strings.Split(step, " ")
		length, _ := strconv.Atoi(parts[1])
		movements = append(movements, Movement{parts[0], length})
	}

	return movements
}

func Locate(steps []string) int {
	horizontal := 0
	depth := 0
	movements := stepsToMovements(steps)

	for _, movement := range movements {
		switch movement.direction {
		case "up":
			depth -= movement.length
		case "down":
			depth += movement.length
		default:
			horizontal += movement.length
		}
	}

	return horizontal * depth
}
