package main

import "fmt"

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i] == 9 {
			if i != 0 {
				digits[i] = 0
				if digits[i-1] < 9 {
					digits[i-1] = digits[i-1] + 1
					return digits
				} else {
					continue
				}
			} else {
				var result = make([]int, len(digits)+1)
				result[0] = 1
				result[i+1] = 0
				return result
			}
		}

		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}

	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
