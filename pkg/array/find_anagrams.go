package main

import (
	"fmt"
	"maps"
	"time"
)

func findAnagrams(s string, p string) []int {
	var result []int

	if len(s) < len(p) {
		return result
	}

	pMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	sMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		sMap[s[i]]++
	}

	if maps.Equal(pMap, sMap) {
		result = append(result, 0)
	}
	right := len(p)
	for i := 1; i <= len(s)-right; i++ {
		if sMap[s[i-1]] == 1 {
			delete(sMap, s[i-1])
		} else {
			sMap[s[i-1]]--
		}
		sMap[s[i+right-1]]++

		if maps.Equal(pMap, sMap) {
			result = append(result, i)
		}
	}

	return result
}
func main() {
	s := "abab"
	p := "ab"
	start := time.Now()
	result := findAnagrams(s, p)
	fmt.Println(result, "Time taken:", time.Since(start))
}
