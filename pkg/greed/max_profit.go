package greed

import "math"

// NO.121
func maxProfit(prices []int) int {

	minPrice := prices[0]
	profit := math.MinInt32
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		profit = max(profit, prices[i]-minPrice)
	}
	return profit
}
