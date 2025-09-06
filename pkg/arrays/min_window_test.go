package arrays

import "testing"

func TestMinWindow(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := minWindow("ADOBECODEBANC", "ABC")
		if v != "BANC" {
			t.Errorf("expected BANC, got %s", v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := minWindow("abc", "ac")
		if v != "abc" {
			t.Errorf("expected abc, got %s", v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd")
		if v != "abbbbbcd" {
			t.Errorf("expected abbbbbcd, got %s", v)
		}
	})

	t.Run("test4", func(t *testing.T) {
		v := minWindow("bbaa", "aba")
		if v != "baa" {
			t.Errorf("expected baa, got %s", v)
		}
	})
}
