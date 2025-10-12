package main

import (
	"fmt"
)

func main() {
	result := (1534236469)

	a := 1<<31 - 1
	fmt.Println(a)
	println(result)
}

func (x int) int {

	var arr []int
	flag := false
	index := 0
	if x < 0 {
		x = -x
		flag = true
	}

	for x >= 10 {
		b := x % 10
		x = x / 10
		if b == 0 && index == 0 {
			continue
		}

		arr = append(arr, b)
		index++
	}

	arr = append(arr, x)

	n, result := len(arr), 0
	for i := 0; i < n; i++ {
		if arr[0] == 0 {
			continue
		}

		result = result*10 + arr[i]
	}

	if flag {
		result = -result
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
