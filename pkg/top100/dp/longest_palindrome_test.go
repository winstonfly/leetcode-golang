package dp

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := longestPalindrome("babad")
		want := "bab"
		if v != want || v == "aba" {
			t.Errorf("longestPalindrome() = %s, want %s", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := longestPalindrome("aaaaa")
		want := "aaaaa"
		if v != want {
			t.Errorf("longestPalindrome() = %s, want %s", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := longestPalindrome("xaabacxcabaaxcabaax")
		want := "xaabacxcabaax"
		if v != want {
			t.Errorf("longestPalindrome() = %s, want %s", v, want)
		}
	})

}
