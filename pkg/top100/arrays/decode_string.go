package arrays

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var backtrack func(string) string
	backtrack = func(s string) string {
		if s == "" {
			return ""
		}

		if !strings.Contains(s, "[") || !strings.Contains(s, "]") {
			return s
		}

		var result []string

		var left, leftIndex, right, rightIndex int
		for i := 0; i < len(s); i++ {
			if s[i] != ']' && s[i] != '[' {
				if leftIndex == 0 && !(s[i] >= '0' && s[i] <= '9') {
					result = append(result, string(s[i]))
				}
				continue
			}
			if s[i] == '[' {
				left++
				if leftIndex == 0 {
					leftIndex = i //记录左括号的索引
				}
			}

			if s[i] == ']' {
				right++
				if right == left {
					rightIndex = i //记录右括号的索引

					//找到匹配的成对的括号
					var (
						count int
						index int
					)
					index = leftIndex
					for ; index > 0; index-- {
						tmpCount, err := strconv.Atoi(s[index-1 : leftIndex])
						if err == nil {
							count = tmpCount
						} else {
							break
						}
					}

					if count == 0 {
						return s // 如果没有数字，直接返回
					}

					// 递归调用
					decoded := backtrack(s[leftIndex+1 : rightIndex])
					// 将解码后的字符串重复 count 次, 同时将不是字符串的部分也添加到结果中

					//result = append(result, s[preIndex:index])
					for ; count > 0; count-- {
						result = append(result, decoded)
					}

					//回溯
					left = 0
					right = 0
					leftIndex = 0
					rightIndex = 0
				}
			}
		}

		return strings.Join(result, "")
	}
	return backtrack(s)
}
