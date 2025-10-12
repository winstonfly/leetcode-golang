package arrays

import (
	"sort"
)

// NO.49 暴力解法
func groupAnagrams2(strs []string) [][]string {
	//用来记录已经分类过
	m := make(map[string]bool, 1)
	var ans [][]string

	for i := 0; i < len(strs); i++ {
		if m[strs[i]] {
			continue
		}
		runes := []rune(strs[i])
		var path []string
		path = append(path, strs[i])
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		for j := i + 1; j < len(strs); j++ {
			runes2 := []rune(strs[j])
			sort.Slice(runes2, func(i, j int) bool {
				return runes2[i] < runes2[j]
			})

			if string(runes) == string(runes2) {
				path = append(path, strs[j])
				m[strs[j]] = true
			}
		}

		ans = append(ans, path)
	}

	return ans
}

func groupAnagrams(strs []string) [][]string {
	//用来记录已经分类过
	m := make(map[string][]string, 1)
	var ans [][]string

	for i := 0; i < len(strs); i++ {
		key := []rune(strs[i])
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})

		if _, ok := m[string(key)]; !ok {
			m[string(key)] = []string{strs[i]}
			continue
		} else {
			m[string(key)] = append(m[string(key)], strs[i])
		}
	}

	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
