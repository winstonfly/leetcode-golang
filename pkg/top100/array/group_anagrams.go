package main

import (
	"maps"
	"slices"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string, 1)
	for i := 0; i < len(strs); i++ {
		runes := []rune(strs[i])
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)

		m[key] = append(m[key], strs[i])
	}

	return slices.Collect(maps.Values(m))
}
