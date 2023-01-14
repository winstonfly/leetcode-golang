package main

import "fmt"

//元素移动法
func removeDuplicates(nums []int) int {

	k := len(nums)
	value := nums[0]
	for i := 1; i <= k-1; {
		if nums[i] == value {
			k--
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
		} else {
			value = nums[i]
			i++
		}
	}

	return k
}

func removeDuplicatesByDoublePointer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for {
		if fast < len(nums) {
			if nums[slow] != nums[fast] {
				slow++
				nums[slow] = nums[fast]
			}
		} else {
			break
		}

		fast++
	}

	return slow + 1
}

func main() {
	fmt.Println(removeDuplicatesByDoublePointer([]int{0, 0, 1, 1, 2}))
}
