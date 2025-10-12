package main

func main() {
	s := "AAAAAAAAAAA"
	result := findRepeatedDnaSequences(s)
	for _, seq := range result {
		println(seq)
	}
}

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int, 1)
	left, right := 0, 10
	var result []string
	for right <= len(s) {
		if index, found := m[s[left:right]]; found {
			if index == 0 {
				result = append(result, s[left:right])
			}
			m[s[left:right]]++ // Mark as found
		} else {
			m[s[left:right]] = 0 // Mark as not found
		}

		left++
		right++
	}

	return result
}
