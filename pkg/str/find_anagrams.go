package str

// NO.438 使用map
func findAnagrams1(s string, p string) []int {
	var ans []int
	length := len(p)
	left, right := 0, 0
	mp := make(map[byte]int, 1)
	for i := 0; i < len(p); i++ {
		mp[p[i]]++
	}

	check := func(m1, m2 map[byte]int) bool {
		for k, v := range m1 {
			if v2, ok := m2[k]; !ok || v2 != v {
				return false
			}
		}

		return true
	}

	ms := make(map[byte]int, 1)
	for right < len(s) {
		if right-left+1 <= length {
			ms[s[right]]++
		}

		if right-left+1 == length {
			if check(ms, mp) {
				ans = append(ans, left)
			}
			ms[s[left]]--
			if ms[s[left]] == 0 {
				delete(ms, s[left])
			}
			left++
		}
		right++
	}

	return ans
}

// NO.438 使用数组
func findAnagrams(s string, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}
	scount, pcount := [26]int{}, [26]int{}
	for i := 0; i < pLen; i++ {
		scount[s[i]-'a']++
		pcount[p[i]-'a']++
	}

	if scount == pcount {
		ans = append(ans, 0)
	}

	for i, v := range s[:sLen-pLen] {
		scount[v-'a']--
		scount[s[i+pLen]-'a']++
		if scount == pcount {
			ans = append(ans, i+1)
		}
	}
	return ans
}
