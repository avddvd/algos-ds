package interviewcake

func findOverLap(point1, length1, point2, length2 int) (int, int) {
	highestStartingPoint := max(point1, point2)
	lowestEndingPoint := min(point1+length1, point2+length2)
	if highestStartingPoint >= lowestEndingPoint {
		return 0, 0
	}
	return highestStartingPoint, (lowestEndingPoint - highestStartingPoint)
}

func findIntersection(rect1, rect2 map[string]int) map[string]int {
	xAxisLowerBound, width := findOverLap(
		rect1["xAxisLowerBound"],
		rect1["width"],
		rect2["xAxisLowerBound"],
		rect2["width"])

	yAxisLowerBound, height := findOverLap(
		rect1["yAxisLowerBound"],
		rect1["height"],
		rect2["yAxisLowerBound"],
		rect2["height"])

	if width == 0 || height == 0 {
		return map[string]int{}
	}
	return map[string]int{
		"xAxisLowerBound": xAxisLowerBound,
		"width":           width,
		"yAxisLowerBound": yAxisLowerBound,
		"height":          height,
	}
}
