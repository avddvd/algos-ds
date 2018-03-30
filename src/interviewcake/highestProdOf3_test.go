package interviewcake

import "testing"

func TestHighestProdOf3(t *testing.T) {
	nums := []int{-10, -10, 1, 3, 2, 4}
	if maxProductOfThree(nums) != 400 {
		t.Error("error calculating highest")
	}
}
