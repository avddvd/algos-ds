package interviewcake

import "testing"

func TestFindIntersection(t *testing.T) {
	rect1 := map[string]int{
		"xAxisLowerBound": 1,
		"width":           6,
		"yAxisLowerBound": 1,
		"height":          3,
	}

	rect2 := map[string]int{
		"xAxisLowerBound": 5,
		"width":           3,
		"yAxisLowerBound": 2,
		"height":          6,
	}

	interSection := findIntersection(rect1, rect2)
	if interSection["xAxisLowerBound"] != 5 ||
		interSection["yAxisLowerBound"] != 2 ||
		interSection["width"] != 2 ||
		interSection["height"] != 2 {
		t.Error("error finding intersection")
	}
}
