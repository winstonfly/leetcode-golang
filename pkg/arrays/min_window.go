package arrays

// NO.76
func minWindow(s string, t string) string {
	tLen := len(t)
	sLen := len(s)

	if sLen < tLen {
		return ""
	}

	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := 0; i < tLen; i++ {
		mt[t[i]]++
	}

	//如果小于，继续向右移动
	check := func(ms, mt map[byte]int) bool {
		for k, v := range mt {
			if ms[k] < v {
				return false
			}
		}
		return true
	}

	ansL, ansR := 0, 0
	var ans string
	for ansR < sLen {
		ms[s[ansR]]++
		for ansR-ansL+1 >= tLen && check(ms, mt) {
			tmp := s[ansL : ansR+1]
			if ans == "" {
				ans = tmp
			} else {
				//更新答案
				if len(ans) > len(tmp) {
					ans = tmp
				}
			}
			//收缩左边界
			ms[s[ansL]]--
			ansL++
		}
		ansR++
	}

	if ansL == 0 && ansR == sLen {
		return ""
	}

	return ans
}
