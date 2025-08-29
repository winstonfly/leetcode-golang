package arrays

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		ans := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
		if ans[0] != 3 || ans[1] != 3 || ans[2] != 5 || ans[3] != 5 || ans[4] != 6 || ans[5] != 7 {
			t.Errorf("expected [3, 3, 5, 5, 6, 7], got %v", ans)
		}
	})

	t.Run("test2", func(t *testing.T) {
		ans := maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5)
		if ans[0] != 9 || ans[1] != 10 || ans[2] != 9 || ans[3] != 9 || ans[4] != 2 {
			t.Errorf("expected [9, 10, 9, 9, 2], got %v", ans)
		}
	})
}
