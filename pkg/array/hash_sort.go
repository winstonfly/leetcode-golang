package main

import (
	"fmt"
	sort2 "leetcode-golang/pkg/sort"
)

func main() {
	arr := []int{1, 5, 2, 3, 1, 4}
	result := sort2.HashSort(arr)
	fmt.Println(result)
}
