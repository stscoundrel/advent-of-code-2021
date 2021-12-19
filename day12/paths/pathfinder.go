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

func canVisitSmallCave(segment string, paths []string) bool {
	canVisitSomeCaveTwice := true
	visits := map[string]int{}

	for _, path := range paths {
		if isSmallCave(path) {
			visits[path] += 1
		}
	}

	for _, count := range visits {
		if count == 2 {
			canVisitSomeCaveTwice = false
		}
	}

	if canVisitSomeCaveTwice {
		return visits[segment] < 2
	}

	return visits[segment] < 1

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

func findDualVisitPaths(currentSegment string, accumulatedPath []string, count int, paths map[string][]string, smallCaveVisits map[string]int) int {
	for _, pathSegment := range paths[currentSegment] {
		if isEnd(pathSegment) {
			count += 1
			continue
		}

		if isStart(pathSegment) {
			continue
		}

		if isSmallCave(pathSegment) {
			if !canVisitSmallCave(pathSegment, append(accumulatedPath, currentSegment)) {
				continue
			}
		}

		currentPathAccumulation := append(accumulatedPath, currentSegment)
		count = findDualVisitPaths(pathSegment, currentPathAccumulation, count, paths, smallCaveVisits)
	}

	return count
}

func CountDualVisitPaths(paths map[string][]string) int {
	accumulatedPath := []string{}
	startSegment := "start"
	smallCaveVisits := map[string]int{}

	return findDualVisitPaths(startSegment, accumulatedPath, 0, paths, smallCaveVisits)
}
