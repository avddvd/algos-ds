package recursion

func PalindromeChecker(src string) bool {
	if len(src) <= 1 {
		return true
	} else {
		return src[0] == src[len(src)-1] && PalindromeChecker(src[1:len(src)-1])
	}
}

func CoinChange(change int, coinVals []int) ([]int, []int) {
	usedCoins := make([]int, change+1)
	minCoins := make([]int, change+1)
	minCoins[0] = 0
	return CoinChangeHelper(change, coinVals, usedCoins, minCoins)
}

func CoinChangeHelper(change int,
	coinVals, usedCoins, minCoins []int) ([]int, []int) {
	for amount := 1; amount <= change; amount++ {
		currMin := amount
		newCoin := 1
		for _, coin := range coinVals {
			if coin <= amount {
				newMin := minCoins[amount-coin] + 1
				if newMin < currMin {
					newCoin = coin
					currMin = newMin
				}
			}
		}
		minCoins[amount] = currMin
		usedCoins[amount] = newCoin
	}
	return usedCoins, minCoins
}

func GetCoinsUsed(usedCoins []int, change int) []int {
	amount := change
	res := []int{}
	for amount > 0 {
		coinUsed := usedCoins[amount]
		amount = amount - coinUsed
		res = append(res, coinUsed)
	}
	return res
}
