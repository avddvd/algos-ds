package leetcode

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	waterVol, currVol, currHeight, currWidth := 0, 0, 0, 0
	for i < j {
		currWidth = j - i
		if height[i] < height[j] {
			currHeight = height[i]
			i += 1
		} else {
			currHeight = height[j]
			j -= 1
		}
		currVol = currWidth * currHeight
		if currVol > waterVol {
			waterVol = currVol
		}
	}
	return waterVol
}
