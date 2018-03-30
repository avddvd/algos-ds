package interviewcake

func productOfAllOtherNums(nums []int) []int {
	beforeProduct := make([]int, len(nums))
	afterProduct := make([]int, len(nums))
	soFarBeforeProd, soFarAfterProd := 1, 1
	for i := 0; i < len(nums); i++ {
		beforeProduct[i] = soFarBeforeProd
		soFarBeforeProd *= nums[i]
	}
	for j := len(nums) - 1; j >= 0; j-- {
		afterProduct[j] = soFarAfterProd
		soFarAfterProd *= nums[j]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = beforeProduct[i] * afterProduct[i]
	}
	return nums
}
