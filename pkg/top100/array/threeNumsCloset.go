package main

import (
	"fmt"
	"leetcode-golang/pkg/sort"
	"math"
)

/**
给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//使用双指针解决
func threeSumCloset(nums []int, target int) int {

	//排序
	nums = sort.Sort(nums)
	best := 1e7
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if len(nums) == 3 {
			return nums[0] + nums[1] + nums[2]
		}
		L := i + 1
		R := len(nums) - 1

		for {
			if L < R {
				value := nums[i] + nums[L] + nums[R]
				if value == target {
					return value
				}

				if math.Abs(float64(value-target)) < math.Abs(best-float64(target)) {
					best = float64(value)
				}

				if value > target {

					for {
						R--
						if R > 0 && L < R && nums[R] == nums[R+1] {
							continue
						} else {
							break
						}
					}
				} else {

					for {
						L++
						if L > 0 && L < R && nums[L] == nums[L-1] {
							continue
						} else {
							break
						}
					}

				}

			} else {
				break
			}

		}
	}

	return int(best)
}

func main() {
	nums := []int{833, 736, 953, -584, -448, 207, 128, -445, 126, 248, 871, 860, 333, -899, 463, 488, -50, -331, 903, 575, 265, 162, -733, 648, 678, 549, 579, -172, -897, 562, -503, -508, 858, 259, -347, -162, -505, -694, 300, -40, -147, 383, -221, -28, -699, 36, -229, 960, 317, -585, 879, 406, 2, 409, -393, -934, 67, 71, -312, 787, 161, 514, 865, 60, 555, 843, -725, -966, -352, 862, 821, 803, -835, -635, 476, -704, -78, 393, 212, 767, -833, 543, 923, -993, 274, -839, 389, 447, 741, 999, -87, 599, -349, -515, -553, -14, -421, -294, -204, -713, 497, 168, 337, -345, -948, 145, 625, 901, 34, -306, -546, -536, 332, -467, -729, 229, -170, -915, 407, 450, 159, -385, 163, -420, 58, 869, 308, -494, 367, -33, 205, -823, -869, 478, -238, -375, 352, 113, -741, -970, -990, 802, -173, -977, 464, -801, -408, -77, 694, -58, -796, -599, -918, 643, -651, -555, 864, -274, 534, 211, -910, 815, -102, 24, -461, -146}
	result := threeSumCloset(nums, -7111)
	fmt.Println(result)

}
