package dp

import "testing"

func TestCoinChange(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		coins := []int{1, 2, 5}
		amount := 11
		v := coinChange(coins, amount)
		want := 3
		if v != want {
			t.Errorf("coinChange() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		v := coinChange(coins, amount)
		want := -1
		if v != want {
			t.Errorf("coinChange() = %d, want %d", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		coins := []int{1}
		amount := 0
		v := coinChange(coins, amount)
		want := 0
		if v != want {
			t.Errorf("coinChange() = %d, want %d", v, want)
		}
	})
}
