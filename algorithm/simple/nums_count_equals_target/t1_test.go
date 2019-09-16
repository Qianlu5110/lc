/*
#
# @Time : 2019/9/16 17:50
# @Author : Qian Lu
*/

package nums_count_equals_target

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	arrIndexs := TwoSum(nums, target)
	fmt.Println(arrIndexs)
}
