package main

import (
	"fmt"
	"leetcode-golang/pkg/sort"
)

//题号88： 合并两个有序的数组
func merge(nums1 []int, m int, nums2 []int, n int) {

	//var nums []int
	//
	//for i := 0; i < m; i++ {
	//	if nums1[i] != 0 {
	//		nums = append(nums, nums1[i])
	//	}
	//}
	//
	//for i := 0; i < n; i++ {
	//	if nums2[i] != 0 {
	//		nums = append(nums, nums2[i])
	//	}
	//}
	//
	//nums1 = nums

	for i := 0; i < n; i++ {
		if nums2[i] != 0 {
			nums1[m+i] = nums2[i]
		}
	}

	//还可以使用copy
	//copy(nums1[m:], nums2)
	sort.Sort(nums1)

	fmt.Println(nums1)

}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5)
}
