package paths

import (
	"strings"
)

func isEnd(segment string) bool {
	return segment == "end"
}

func isStart(segment string) bool {
	return segment == "start"
}

func isSmallCave(segment string) bool {
	return segment == strings.ToLower(segment)
}

func isVisited(segment string, paths []string) bool {
	for _, path := range paths {
		if path == segment {
			return true
		}
	}

	return false
}

func findPaths(currentSegment string, accumulatedPath []string, count int, paths map[string][]string) int {
	for _, pathSegment := range paths[currentSegment] {
		if isEnd(pathSegment) {
			count += 1
			continue
		}

		if isStart(pathSegment) {
			continue
		}

		if isSmallCave(pathSegment) && isVisited(pathSegment, accumulatedPath) {
			continue
		}

		currentPathAccumulation := append(accumulatedPath, currentSegment)
		count = findPaths(pathSegment, currentPathAccumulation, count, paths)
	}

	return count
}

func CountPaths(paths map[string][]string) int {
	accumulatedPath := []string{}
	startSegment := "start"

	return findPaths(startSegment, accumulatedPath, 0, paths)
}
