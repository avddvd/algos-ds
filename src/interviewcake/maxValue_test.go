package interviewcake

import "testing"

func TestMaxValue(t *testing.T) {
	input := [][]int{[]int{7, 160}, []int{3, 90}, []int{2, 15}}
	result := maxValue(input, 20)
	if result != 555 {
		t.Error("error calculating max value", result)
	}
}
