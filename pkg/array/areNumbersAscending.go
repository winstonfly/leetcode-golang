package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello world 5 x 5"
	fmt.Println(areNumbersAscending(s))
}

func areNumbersAscending(s string) bool {
	arr := strings.Split(s, " ")
	n := len(arr)
	prev := 0
	flag := true
	for i := 0; i < n; i++ {
		a, err := strconv.Atoi(arr[i])
		if err != nil {
			continue
		}

		if a > prev {
			prev = a
			continue
		}

		flag = false
	}

	return flag
}
