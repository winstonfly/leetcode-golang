package stacks

// NO.20
func isValid(s string) bool {

	var stack []byte

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}

		switch c {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
