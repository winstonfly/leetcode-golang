package dp

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := longestPalindrome("babad")
		t.Log(v)
	})

	t.Run("test2", func(t *testing.T) {
		v := longestPalindrome("aaaaa")
		t.Log(v)

	})
}

func TestPartition(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := partition("aab")
		t.Log(v)
	})
}
