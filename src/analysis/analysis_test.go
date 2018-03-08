package analysis

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{100, 113, 110, 85, 105, 102, 86, 63, 81,
		101, 94, 106, 101, 79, 94, 90, 97}
	buy, sell, maxProfit := MaxProfit(prices)
	fmt.Println(buy, sell, maxProfit)
	prices = []int{10, 7, 5, 8, 11, 9}
	fmt.Println(MaxProfit(prices))
	prices = []int{10, 9, 7, 6, 5}
	fmt.Println(MaxProfit(prices))
}
