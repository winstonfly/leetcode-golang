package main

import (
	"fmt"
)

//No. 136
//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

//你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/single-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func singleNumber(nums []int) int {

	var value int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = m[nums[i]] + 1
	}
	for k, v := range m {
		if v == 1 {
			value = k
			break
		}
	}

	return value
}

//XOR：位运算解题
func singleNumberByXOR(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

//No.137
func singleNumber3(nums []int) int {
	var value int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = m[nums[i]] + 1
	}
	for k, v := range m {
		if v == 1 {
			value = k
			break
		}
	}

	return value
}

func main() {
	fmt.Println(singleNumber3([]int{2, 2, 3, 2}))
}
