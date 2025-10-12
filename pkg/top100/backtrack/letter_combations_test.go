package backtrack

import (
	"slices"
	"testing"
)

func TestLetterCombations(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := "23"
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		if got := letterCombinations(tests); !slices.Equal(got, want) {
			t.Errorf("letterCombinations() = %v, want %v", got, want)
		}
	})
}
