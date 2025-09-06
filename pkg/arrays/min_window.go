package arrays

import "math"

// NO.76
func minWindow(s string, t string) string {
	ms, mt := make(map[byte]int, 1), make(map[byte]int, 1)
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}

	sLen := math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range mt {
			for ms[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && ms[s[r]] > 0 {
			ms[s[r]]++
		}

		for check() && l <= r {
			if r-l+1 < sLen {
				sLen = r - l + 1
				ansL, ansR = l, r
			}

			if ms[s[l]] > 0 {
				ms[s[l]]--
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}
	return s[ansL : ansR+1]
}
