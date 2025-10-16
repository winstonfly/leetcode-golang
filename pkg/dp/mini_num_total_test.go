package dp

import "testing"

func TestMiniMumTotal(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := minimumTotal([][]int{
			{-10},
		})
		want := -10
		if v != want {
			t.Errorf("minimumTotal() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := minimumTotal([][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		})
		want := 11
		if v != want {
			t.Errorf("minimumTotal() = %d, want %d", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := minimumTotal([][]int{
			{-1},
			{2, 3},
			{1, -1, -3},
		})
		want := -1
		if v != want {
			t.Errorf("minimumTotal() = %d, want %d", v, want)
		}
	})

	t.Run("test4", func(t *testing.T) {
		v := minimumTotal([][]int{
			{1},
			{2, 3},
		})
		want := 3
		if v != want {
			t.Errorf("minimumTotal() = %d, want %d", v, want)
		}
	})
}
