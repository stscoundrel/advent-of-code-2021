package coordinates

func isVertical(line Line) bool {
	return line.start.x == line.end.x
}

func isHorizontal(line Line) bool {
	return line.start.y == line.end.y
}

func isVerticalOrHorizontal(line Line) bool {
	return isVertical(line) || isHorizontal(line)
}

func getStartAndEndCoordinates(line Line, vertical bool) (Coordinate, Coordinate) {
	if vertical {
		if line.start.y < line.end.y {
			return line.start, line.end
		} else {
			return line.end, line.start
		}
	}

	if line.start.x < line.end.x {
		return line.start, line.end
	}

	return line.end, line.start
}

func getStartAndEndPoints(start Coordinate, end Coordinate, vertical bool) (int, int) {
	if vertical {
		return start.y, end.y
	}

	return start.x, end.x
}

func interpolatePoints(line Line, vertical bool) (linePoints []Coordinate) {
	startCoordinate, endCoordinate := getStartAndEndCoordinates(line, vertical)
	startPoint, endPoint := getStartAndEndPoints(startCoordinate, endCoordinate, vertical)

	linePoints = append(linePoints, startCoordinate)

	for startPoint < endPoint {
		startPoint++
		newPoint := line.start
		if vertical {
			newPoint.y = startPoint
		} else {
			newPoint.x = startPoint
		}

		linePoints = append(linePoints, newPoint)
	}

	return
}

func interpolateDiagonalPoints(line Line) (linePoints []Coordinate) {
	startCoordinate, endCoordinate := getStartAndEndCoordinates(line, false)

	linePoints = append(linePoints, startCoordinate)

	for startCoordinate.x < endCoordinate.x {
		startCoordinate.x = startCoordinate.x + 1

		if startCoordinate.y > endCoordinate.y {
			startCoordinate.y = startCoordinate.y - 1
		} else {
			startCoordinate.y = startCoordinate.y + 1
		}

		newPoint := startCoordinate

		linePoints = append(linePoints, newPoint)
	}

	return
}

func GetPointsOfLine(line Line) (coordinates []Coordinate) {
	if isVerticalOrHorizontal(line) {
		coordinates = interpolatePoints(line, isVertical((line)))
	} else {
		coordinates = interpolateDiagonalPoints(line)
	}

	return
}
