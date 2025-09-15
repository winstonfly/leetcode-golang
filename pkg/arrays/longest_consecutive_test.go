package arrays

import "testing"

func TestLongestConsecutive(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
		t.Log(v)
	})

	t.Run("test1", func(t *testing.T) {
		v := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
		t.Log(v)
	})

	t.Run("test2", func(t *testing.T) {
		v := longestConsecutive([]int{0, -1})
		t.Log(v)
	})

	t.Run("test3", func(t *testing.T) {
		v := longestConsecutive([]int{100, 4, 200, 1})
		t.Log(v)
	})

	t.Run("test4", func(t *testing.T) {
		v := longestConsecutive([]int{0, 1, 2, 4, 8, 5, 6, 7, 9, 3, 55, 88, 77, 99, 999999999})
		t.Log(v)
	})
}
