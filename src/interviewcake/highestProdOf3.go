package interviewcake

func max3(int1, int2, int3 int) int {
	maxNum := max(int1, int2)
	maxNum = max(maxNum, int3)
	return maxNum
}

func max(int1, int2 int) int {
	var maxNum int
	if int1 > int2 {
		maxNum = int1
	} else {
		maxNum = int2
	}
	return maxNum
}

func min(int1, int2 int) int {
	var minNum int
	if int1 > int2 {
		minNum = int2
	} else {
		minNum = int1
	}
	return minNum
}

func min3(int1, int2, int3 int) int {
	minNum := min(int1, int2)
	minNum = min(minNum, int3)
	return minNum
}

func maxProductOfThree(nums []int) int {
	maxProd2 := nums[0] * nums[1]
	minProd2 := nums[0] * nums[1]
	maxProd3 := nums[0] * nums[1] * nums[2]
	highest := max(nums[0], nums[1])
	lowest := min(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		current := nums[i]
		maxProd3 = max3(maxProd3, maxProd2*current, minProd2*current)
		maxProd2 = max3(maxProd2, highest*current, lowest*current)
		minProd2 = min3(minProd2, highest*current, lowest*current)
		highest = max(highest, current)
		lowest = min(highest, current)
	}
	return maxProd3
}
