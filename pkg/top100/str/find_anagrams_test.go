package str

import (
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		s := "cbaebabacd"
		p := "abc"
		expected := []int{0, 6}
		if got := findAnagrams(s, p); !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test2", func(t *testing.T) {
		s := "abab"
		p := "ab"
		expected := []int{0, 1, 2}
		if got := findAnagrams(s, p); !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test3", func(t *testing.T) {
		s := "af"
		p := "be"
		expected := []int{}
		if got := findAnagrams(s, p); !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test4", func(t *testing.T) {
		s := "aaaaaaaaaa"
		p := "aaaaaaaaaaaaa"
		expected := []int{}
		if got := findAnagrams(s, p); !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
