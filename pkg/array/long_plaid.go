package main

import (
	"fmt"
	"strings"
)

func isPalindrome(x int) bool {

	var s string

	if x > 0 && x < 10 {
		return true
	} else if x < 0 {
		return false
	}
	s = fmt.Sprintf("%d", x)
	n := len(s)

	flag := true
	i := 0
	j := n - 1
	for i < j {
		if s[i] != s[j] {
			flag = false
			break
		}
		i++
		j--
	}

	return flag
}

func intToRoman(num int) string {

	var intToRomanFn func(num int, result []string) []string

	intToRomanFn = func(num int, result []string) []string {
		if num > 1000 {
			t := num / 1000
			for i := 0; i < t; i++ {
				result = append(result, "M")
			}
			return intToRomanFn(num-t*1000, result)
		} else if num > 100 {
			t := num / 100
			switch t {
			case 4:
				result = append(result, "CD")
			case 9:
				result = append(result, "CM")
			default:
				if t < 5 {
					for i := 0; i < t; i++ {
						result = append(result, "C")
					}
				} else {
					result = append(result, "D")
					for i := 0; i < t-5; i++ {
						result = append(result, "C")
					}
				}

			}
			return intToRomanFn(num-t*100, result)
		} else if num > 10 {
			t := num / 10
			switch t {
			case 4:
				result = append(result, "XL")
			case 9:
				result = append(result, "XC")
			default:
				if t < 5 {
					for i := 0; i < t; i++ {
						result = append(result, "X")
					}
				} else {
					result = append(result, "L")
					for i := 0; i < t-5; i++ {
						result = append(result, "X")
					}
				}
			}
			return intToRomanFn(num-t*10, result)
		} else {
			switch num {
			case 4:
				result = append(result, "IV")
			case 9:
				result = append(result, "IX")
			default:
				if num > 5 {
					t := num - 5
					result = append(result, "V")
					for i := 0; i <= t; i++ {
						result = append(result, "I")
					}
				} else {
					t := num
					for i := 0; i <= t; i++ {
						result = append(result, "I")
					}
				}

			}

		}

		return result
	}

	var tmp []string
	result := intToRomanFn(num, tmp)
	return strings.Join(result, "")
}

func main() {
	fmt.Println(isPalindrome(1001))
}
