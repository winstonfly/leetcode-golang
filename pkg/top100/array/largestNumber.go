package main

import (
	"fmt"
	"leetcode-golang/pkg/sort"
)

//No.179 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

//注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

func largestNumber(nums []int) string {

	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		v1 := nums[i] / 10
		if m[v1] == nil {
			var arr []int
			m[v1] = arr

		}
		m[v1] = append(m[v1], nums[i])
	}

	var result string
	for i := 9; i > 0; i-- {
		arrs := m[i]
		if arrs != nil {
			arrs = sort.Sort(arrs)
			result = fmt.Sprintf("%s%s", result, arrs)
		}
	}

	return result
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
}
