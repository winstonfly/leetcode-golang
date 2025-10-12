package arrays

import "math"

// NO.121
func maxProfit1(prices []int) int {
	ans := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] < prices[i] {
				continue
			}
			ans = max(ans, prices[j]-prices[i])
		}
	}

	return ans
}

func maxProfit(prices []int) int {
	minn, maxprofit := math.MaxInt32, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minn {
			minn = prices[i]
		} else if prices[i]-minn > maxprofit {
			maxprofit = prices[i] - minn
		}
	}

	return maxprofit
}
