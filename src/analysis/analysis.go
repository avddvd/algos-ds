package analysis

func SumToN(array []int) int {
	var output int
	for _, val := range array {
		output += val
	}
	return output
}

func Anagram(src, dst string) bool {
	if len(src) != len(dst) {
		return false
	}
	if src == dst {
		return true
	}
	counts := make(map[rune]int)
	for _, val := range src {
		if _, ok := counts[val]; ok {
			counts[val] = counts[val] + 1
		} else {
			counts[val] = 1
		}
	}

	for _, val := range dst {
		if _, ok := counts[val]; ok {
			counts[val] = counts[val] - 1
		} else {
			return false
		}
	}

	for _, val := range counts {
		if val != 0 {
			return false
		}
	}
	return true
}

func MaxProfit(prices []int) (int, int, int) {
	buy, sell := 0, 0
	maxProfit := -1
	i := 0
	for i < len(prices) {
		if prices[i] < prices[buy] {
			buy = i
		}
		if prices[i]-prices[buy] > maxProfit {
			maxProfit = prices[i] - prices[buy]
			sell = i
		}
		i += 1
	}
	return buy, sell, maxProfit
}
