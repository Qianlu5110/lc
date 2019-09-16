/*
#
# @Time : 2019/9/16 17:45
# @Author : Qian Lu
*/

package nums_count_equals_target

import (
	"fmt"
	"sort"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TwoSum(nums []int, target int) []int {
	backOff := make([]int, len(nums), len(nums))
	copy(backOff, nums)
	sort.Ints(backOff)

	var var1, var2 int
	for i := 0; i < len(backOff); i++ {
		if backOff[i] <= target {
			for j := 0; j <= len(backOff); j++ {
				if i == j || backOff[j] > target {
					break
				}
				if backOff[i]+backOff[j] == target {
					var1 = backOff[i]
					var2 = backOff[j]
				}
			}
		} else {
			break
		}
	}

	var index []int
	if var1 != 0 && var2 != 0 {
		for n := 0; n < len(nums); n++ {
			if nums[n] == var1 || nums[n] == var2 {
				fmt.Println(n)
				fmt.Println(nums[n])
				index = append(index, n)
			}
		}
	}

	return index
}
