package dp

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := longestPalindromeSubseq("bbbab")
		want := 4
		if v != want {
			t.Errorf("longestPalindromeSubseq() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := longestPalindromeSubseq("cbbd")
		want := 2
		if v != want {
			t.Errorf("longestPalindromeSubseq() = %d, want %d", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := longestPalindromeSubseq("abcabcabcabc")
		want := 7
		if v != want {
			t.Errorf("longestPalindromeSubseq() = %d, want %d", v, want)
		}
	})
}

/**
dp[i:j]代表s[i:j的最大回文子序列的长度
以abcabcabcabc为例:
dp[0][0] = 1
dp[0][1] = 1
dp[0][2] = 1
dp[0][3] = 3

*/
