package arrays

import (
	"slices"
	"testing"
)

func TestMergeArray(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		ans := merge(intervals)
		expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
		if slices.EqualFunc(ans, expected, func(ints []int, ints2 []int) bool {
			return ints[0] == ints2[0] && ints[1] == ints2[1]
		}) {
			t.Errorf("got %v, want %v", ans, expected)
		}
	})

	t.Run("test2", func(t *testing.T) {
		intervals := [][]int{{1, 4}, {4, 5}}
		ans := merge(intervals)
		expected := [][]int{{1, 5}}
		if slices.EqualFunc(ans, expected, func(ints []int, ints2 []int) bool {
			return ints[0] == ints2[0] && ints[1] == ints2[1]
		}) {
		}
	})
}
