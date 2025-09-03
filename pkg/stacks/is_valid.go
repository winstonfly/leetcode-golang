package stacks

// NO.20
func isValid(s string) bool {

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if top == '(' && c == ')' || top == '[' && c == ']' || top == '{' && c == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	if len(stack) != 0 {
		return false
	}

	return true
}
