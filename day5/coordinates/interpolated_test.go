package coordinates

import (
	"fmt"
	"testing"
)

func TestGetInBetweenPointsOfLineVertical(t *testing.T) {
	line := Line{
		start: Coordinate{
			x: 2,
			y: 0,
		},
		end: Coordinate{
			x: 2,
			y: 5,
		},
	}

	result := GetPointsOfLine(line)

	for index, point := range result {
		if point.x != 2 {
			t.Error("Did not return expected X coordinate. Received", point.x, "expected ", 2)
		}

		if point.y != line.start.y+index {
			t.Error("Did not return expected Y coordinate. Received", point.y, "expected ", line.start.y+index)
		}
	}
}

func TestGetInBetweenPointsOfLineHorizontal(t *testing.T) {
	line := Line{
		start: Coordinate{
			x: 2,
			y: 0,
		},
		end: Coordinate{
			x: 5,
			y: 0,
		},
	}

	result := GetPointsOfLine(line)

	fmt.Println(result)

	for index, point := range result {
		if point.y != 0 {
			t.Error("Did not return expected Y coordinate. Received", point.y, "expected ", 0)
		}

		if point.x != line.start.x+index {
			t.Error("Did not return expected X coordinate. Received", point.x, "expected ", line.start.x+index)
		}
	}
}
