package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {

	var result []int

	n := len(nums)

	preDp := make([]int, len(nums))
	sufDp := make([]int, len(nums))

	preDp[0] = 1
	for i := 1; i < n; i++ {
		preDp[i] = preDp[i-1] * nums[i-1]
	}

	sufDp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		sufDp[i] = sufDp[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result = append(result, preDp[i]*sufDp[i])
	}

	return result

}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
