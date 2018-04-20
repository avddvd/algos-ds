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
