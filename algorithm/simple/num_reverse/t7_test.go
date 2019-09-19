/*
#
# @Time : 2019/9/18 15:00
# @Author : Qian Lu
*/

package num_reverse

import (
	"fmt"
	"testing"
)

func Test_NumReverse(t *testing.T) {
	t1 := Reverse(-123456789)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", -123456789, t1))

	t2 := Reverse(123456789)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", 123456789, t2))

	t3 := Reverse(0)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", 0, t3))

	t4 := Reverse(1)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", 0, t4))

	t5 := Reverse(1563847412)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", 1563847412, t5))

	t6 := Reverse(1534236469)
	fmt.Println(fmt.Sprintf("[reverse]before: %d ,result : %d", 1534236469, t6))

}
