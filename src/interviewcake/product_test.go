package interviewcake

import "testing"

func TestProductOfAllOtherNums(t *testing.T) {
	nums := []int{1, 7, 3, 4}
	expected := []int{84, 12, 28, 21}
	results := productOfAllOtherNums(nums)
	for i := 0; i < len(nums); i++ {
		if results[i] != expected[i] {
			t.Error("productOfAllOtherNums failed", expected, results)
		}
	}
}
