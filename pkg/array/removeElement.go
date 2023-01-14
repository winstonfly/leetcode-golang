package main

import "fmt"

//题号: 27
func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for {
		if fast < len(nums) {
			if nums[fast] != val {
				nums[slow] = nums[fast]
				slow++
			}
		} else {
			break
		}
		fast++
	}

	fmt.Println(nums)
	return slow
}

func main() {

	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
