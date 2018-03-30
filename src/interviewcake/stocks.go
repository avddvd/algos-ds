package interviewcake

func maxProfit(stocks []int) (int, int, int) {
	min_buy := 0
	max_sell := 1
	max_profit := stocks[max_sell] - stocks[min_buy]
	for i := 1; i < len(stocks); i++ {
		curr_price := stocks[i]
		if curr_price-stocks[min_buy] > max_profit {
			max_profit = curr_price - stocks[min_buy]
			max_sell = curr_price
		}
		if curr_price < stocks[min_buy] {
			min_buy = i
		}
	}
	return max_profit, min_buy, max_sell
}
