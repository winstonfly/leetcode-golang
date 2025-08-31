package arrays

import "testing"

func TestMergeArray(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		ans := merge(intervals)
		t.Log(ans)
	})
}
