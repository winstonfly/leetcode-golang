package main

import "strconv"

func main() {
	result := myAtoi("21474836460")
	println(result) // Expected output: -42
}

func myAtoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[0] == ' ' {
			s = s[1:]
			continue
		}
	}

	n := len(s)
	var flag bool
	if n <= 0 {
		return 0
	}

	switch s[0] {
	case '-':
		s = s[1:]
		flag = true
	case '+':
		s = s[1:]
	}

	y := 0
	for _, v := range s {
		a, err := strconv.Atoi(string(v))
		if err != nil {
			break
		}

		y = y*10 + a%10
		if y >= (1 << 31) {
			if !flag {
				y = 1<<31 - 1
			} else {
				y = 1 << 31
			}
			break
		}

	}

	if flag {
		y = -y
	}

	return y

}
