package str

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		s := "abcabcbb"
		expected := 3
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test2", func(t *testing.T) {
		s := "bbbbb"
		expected := 1
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test3", func(t *testing.T) {
		s := "pwwkew"
		expected := 3
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test4", func(t *testing.T) {
		s := ""
		expected := 0
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test5", func(t *testing.T) {
		s := "au"
		expected := 2
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test6", func(t *testing.T) {
		s := "aab"
		expected := 2
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("test7", func(t *testing.T) {
		s := "cdd"
		expected := 2
		if got := lengthOfLongestSubstring(s); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
