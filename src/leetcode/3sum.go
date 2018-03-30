package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				// sum negative, move l to right to find 0 sum
				l += 1
			} else if sum > 0 {
				// sum positive, move r to left to find 0 sum
				r -= 1
			} else if sum == 0 {
				results = append(results, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r = r - 1
				}
				l += 1
				r -= 1
			}
		}
	}
	return results
}
