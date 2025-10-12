package stacks

import (
	"strconv"
	"strings"
)

// NO.394
func decodeString(s string) string {
	var stack []string
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ']' {
			if c >= '0' && c <= '9' {
				//如果是数字，可能是多位数
				var num string
				for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
					num += string(s[i])
				}
				stack = append(stack, num)
				i--
				continue
			}
			stack = append(stack, string(c))
			continue
		}
		//开始出栈
		if c == ']' {
			var str string
			var top string
			var pre string
			top = stack[len(stack)-1]
			for top != "[" {
				stack = stack[:len(stack)-1]
				str = top + str
				top = stack[len(stack)-1]
			}
			//弹出左括号
			stack = stack[:len(stack)-1]
			num := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count, _ := strconv.Atoi(num)
			var tmpStr string
			for j := 0; j < count; j++ {
				tmpStr += str
			}

			stack = append(stack, pre+tmpStr)
		}
	}

	return strings.Join(stack, "")
}

func decodeString1(s string) string {
	var ans string

	//var backtrack func() string
	//backtrack = func() string {
	//	if s == "" {
	//		return ""
	//	}
	//
	//	var curr string
	//	var count int
	//	for i := 0; i < len(s); i++ {
	//		c := s[i]
	//	}
	//
	//}

	return ans
}
