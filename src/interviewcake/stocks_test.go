package interviewcake

import "testing"

func TestStocks(t *testing.T) {
	stocks := []int{10, 7, 5, 8, 11, 9}
	p, b, s := maxProfit(stocks)
	if p != 6 && b != 4 && s != 5 {
		t.Error("error calculating maxProfit:", p, b, s)
	}
}
