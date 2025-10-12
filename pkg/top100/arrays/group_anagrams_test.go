package arrays

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tests := []struct {
			input  []string
			output [][]string
		}{
			{
				input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
				output: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
			},
		}

		v := groupAnagrams(tests[0].input)
		t.Log(v)
	})
}
