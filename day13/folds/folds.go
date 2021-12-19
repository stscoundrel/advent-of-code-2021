package folds

func separateByFold(points []Point, fold Fold) ([]Point, []Point) {
	keep := []Point{}
	transform := []Point{}

	for _, point := range points {
		if fold.direction == "y" && (point.y < fold.value) || fold.direction == "x" && (point.x < fold.value) {
			keep = append(keep, point)
		} else {
			transform = append(transform, point)
		}
	}

	return keep, transform
}

func foldPoints(points []Point, fold Fold) []Point {
	foldedPoints := []Point{}

	for _, point := range points {
		if fold.direction == "y" {
			newY := fold.value - (point.y - fold.value)
			foldedPoints = append(foldedPoints, Point{x: point.x, y: newY})
		} else {
			newX := fold.value - (point.x - fold.value)
			foldedPoints = append(foldedPoints, Point{x: newX, y: point.y})
		}
	}

	return foldedPoints
}

func doFold(points []Point, fold Fold) []Point {
	keep, transform := separateByFold(points, fold)
	transformed := foldPoints(transform, fold)

	return append(keep, transformed...)
}

func CountSpaces(points []Point, folds []Fold) int {
	folded := points

	for _, fold := range folds {
		folded = doFold(folded, fold)
	}

	// for _, point := range folded {
	// 	fmt.Println("X:", point.x, " y:", point.y)
	// }

	return calculateDots(folded)
}
