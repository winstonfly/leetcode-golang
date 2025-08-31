package arrays

import "testing"

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	if result := maxProfit(prices); result != expected {
		t.Errorf("maxProfit(%v) = %d; want %d", prices, result, expected)
	}

	prices = []int{7, 6, 4, 3, 1}
	expected = 0
	if result := maxProfit(prices); result != expected {
		t.Errorf("maxProfit(%v) = %d; want %d", prices, result, expected)
	}
}
