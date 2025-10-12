package main

import "strings"

func longestCommonPrefix(strs []string) string {

	minLen := 0
	var minLenStr string
	for _, v := range strs {
		if minLen == 0 {
			minLen = len(v)
			minLenStr = v
		} else {
			if len(v) < minLen {
				minLen = len(v)
				minLenStr = v
			}
		}
	}

	for i := minLen; i >= 0; i-- {
		flag := true
		for _, v := range strs {
			if !strings.HasPrefix(v, minLenStr[0:i]) {
				flag = false
			}
		}

		if flag {
			return minLenStr[0:i]
		}
	}

	return ""
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	println(result) // Output: "fl"
}
