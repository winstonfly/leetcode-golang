package main

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 1 {
		return
	}

	move := 0

	if k < n {
		move = k
	} else {
		if k/n > 1 {
			move = k - (k/n-1)*(k/n)
		} else {
			move = k - n
		}
	}
	var tmp []int
	tmp = append(tmp, nums[n-move:]...)

	// 1 2 3 4 5 6 7
	copy(nums[move:], nums[0:n-move])
	copy(nums[:move], tmp)
}

func main() {

	nums := []int{1, 2}
	k := 3
	rotate(nums, k)
	for _, v := range nums {
		println(v)
	}
}
