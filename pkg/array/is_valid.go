package main

// No.20
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

//有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if !match(top, c) {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

func match(left, right byte) bool {
	switch left {
	case '(':
		return right == ')'
	case '[':
		return right == ']'
	case '{':
		return right == '}'
	default:
		return false
	}
}

func main() {
	s := "()[]{}"
	result := isValid(s)
	if result {
		println("The string is valid.")
	} else {
		println("The string is not valid.")
	}
}
