package greed

// NO.763
func partitionLabels(s string) []int {
	var ans []int
	//用来存储每个字母最后出现的位置
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}

	start, end := 0, 0
	for i := start; i < len(s); i++ {
		end = max(end, m[s[i]])
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}
