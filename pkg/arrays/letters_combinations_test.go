package arrays

import "testing"

func TestLetterCombinations(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		v := letterCombinations("23")
		t.Log(v)
	})
}
