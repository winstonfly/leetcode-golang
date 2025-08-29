package dp

import "testing"

func TestCanPartition(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{1, 2, 3, 5}
		v := canPartition(nums)
		if v {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("test2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		v := canPartition(nums)
		if !v {
			t.Errorf("expected true, got false")
		}
	})
}
