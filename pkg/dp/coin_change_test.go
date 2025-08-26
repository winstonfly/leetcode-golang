package dp

import "testing"

func TestCoinChange(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		coins := []int{1, 2, 5}
		amount := 11
		v := coinChange(coins, amount)
		t.Log(v)
	})
}
