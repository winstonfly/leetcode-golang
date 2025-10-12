package main

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutive(nums)
	println(result) // Output: 4
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	m := make(map[int]int, 1)
	for _, v := range nums {
		m[v] = v
	}

	result := 1

	for key, _ := range m {
		if _, ok := m[key-1]; ok {
			continue
		} else {
			r := 0
			for {
				if _, ok = m[key]; ok {
					r++
					key++
					result = max(result, r)
				} else {
					break
				}
			}
		}
	}

	return result
}
