package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	t.Run("Test length of LIS", func(t *testing.T) {
		tests := []int{10, 9, 2, 5, 3, 7, 101, 18}
		v := lengthOfLIS(tests)
		want := 4
		if v != want {
			t.Errorf("LengthOfLIS expect %d, got %d", want, v)
		}
	})

	t.Run("Test length of LIS 2", func(t *testing.T) {
		tests := []int{0, 1, 0, 3, 2, 3}
		v := lengthOfLIS(tests)
		want := 4
		if v != want {
			t.Errorf("LengthOfLIS expect %d, got %d", want, v)
		}
	})

	t.Run("Test length of LIS 3", func(t *testing.T) {
		tests := []int{7, 7, 7, 7, 7, 7, 7}
		v := lengthOfLIS(tests)
		want := 1
		if v != want {
			t.Errorf("LengthOfLIS expect %d, got %d", want, v)
		}
	})

	t.Run("Test length of LIS 4", func(t *testing.T) {
		tests := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
		v := lengthOfLIS(tests)
		want := 6
		if v != want {
			t.Errorf("LengthOfLIS expect %d, got %d", want, v)
		}
	})

}
