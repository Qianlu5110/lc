/*
#
# @Time : 2019/9/17 16:13
# @Author : Qian Lu
*/

package nums_count

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 显示所有链表元素
func (l *ListNode) AllNode() {
	str := "ListNode: "
	str = str + fmt.Sprintf("%d", l.Val)
	for l.Next != nil {
		str = str + fmt.Sprintf(" -> %d", l.Next.Val)
		l = l.Next
	}

	str = strings.TrimRight(str, "->")
	fmt.Println(str)
}
