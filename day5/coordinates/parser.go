package coordinates

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Line struct {
	start Coordinate
	end   Coordinate
}

func parseCoordinate(rawX string, rawY string) Coordinate {
	x, _ := strconv.Atoi(rawX)
	y, _ := strconv.Atoi(rawY)

	return Coordinate{x: x, y: y}
}

func parseLine(line string) Line {
	parts := strings.Split(line, " -> ")
	startParts := strings.Split(parts[0], ",")
	endParts := strings.Split(parts[1], ",")

	start := parseCoordinate(startParts[0], startParts[1])
	end := parseCoordinate(endParts[0], endParts[1])

	return Line{start: start, end: end}
}

func ParseCoordinates(rawLines []string) []Line {
	lines := []Line{}

	for _, rawLine := range rawLines {
		line := parseLine(rawLine)
		lines = append(lines, line)
	}

	return lines
}
