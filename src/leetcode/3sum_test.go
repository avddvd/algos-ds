package leetcode

import ("testing")

func Test3Sum(t *testing.T) {
  r := threeSum([]int{0,0,0})
  if len(r) == 0 {
    t.Error("error in binarySearch")
  }
}
