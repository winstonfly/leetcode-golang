package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	t.Run("Test length of LIS", func(t *testing.T) {
		tests := []int{10, 9, 2, 5, 3, 7, 101, 18}
		v := lengthOfLIS(tests)
		if v != 4 {
			t.Errorf("LengthOfLIS expect 4, got %d", v)
		}
	})

	t.Run("Test length of LIS 2", func(t *testing.T) {
		tests := []int{0, 1, 0, 3, 2, 3}
		v := lengthOfLIS(tests)
		if v != 4 {
			t.Errorf("LengthOfLIS expect 4, got %d", v)
		}
	})

	t.Run("Test length of LIS 3", func(t *testing.T) {
		tests := []int{7, 7, 7, 7, 7, 7, 7}
		v := lengthOfLIS(tests)
		if v != 1 {
			t.Errorf("LengthOfLIS expect 1, got %d", v)
		}
	})
}
