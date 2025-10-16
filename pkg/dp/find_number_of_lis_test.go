package dp

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := findNumberOfLIS([]int{1, 3, 5, 4, 7})
		want := 2
		if v != want {
			t.Errorf("findNumberOfLIS() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := findNumberOfLIS([]int{2, 2, 2, 2, 2})
		want := 5
		if v != want {
			t.Errorf("findNumberOfLIS() = %d, want %d", v, want)
		}
	})

}
