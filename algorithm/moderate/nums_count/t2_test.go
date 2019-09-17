/*
#
# @Time : 2019/9/17 15:53
# @Author : Qian Lu
*/

package nums_count

import (
	"testing"
)

func Test_NumsCount(t *testing.T) {

	la := new(ListNode)
	la.Val = 9

	la1 := new(ListNode)
	la1.Val = 9
	la.Next = la1

	lb := new(ListNode)
	lb.Val = 1

	// la.AllNode()
	// lb.AllNode()

	res := addTwoNumbers(la, lb)
	res.AllNode()

}
