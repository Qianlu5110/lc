/*
#
# @Time : 2019/9/17 15:53
# @Author : Qian Lu
*/

package nums_count

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := new(ListNode)
	carry := 0

	leftChild := new(ListNode)
	curr.Next = leftChild
	for l1 != nil || l2 != nil {
		var x = 0
		var y = 0
		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum / 10

		rightChild := new(ListNode)
		leftChild.Val = sum % 10
		leftChild.Next = rightChild
		leftChild = rightChild

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		rightChild := new(ListNode)
		rightChild.Val = carry
		leftChild.Next = rightChild
	}

	return curr.Next
}

// 第一想法，这个不太好
func old(l1 *ListNode, l2 *ListNode) *ListNode {
	var list1Vals []int
	var index = 0
	list1Vals = append(list1Vals, l1.Val)

	for l1.Next != nil {
		list1Vals = append(list1Vals, l1.Next.Val)
		l1 = l1.Next
	}

	res := new(ListNode)
	res.Val = (list1Vals[index] + l2.Val) / 1 % 10
	leftChild := res

	// 定义进位
	carry := 0
	if list1Vals[0]+l2.Val >= 10 {
		carry = 1
	}
	if l2.Next == nil && carry == 1 && len(list1Vals) > 1 {
		index = index + 1
		rightChild := new(ListNode)
		rightChild.Val = carry + list1Vals[index]
		res.Next = rightChild
	}

	for l2.Next != nil {
		index++
		l1Val := 0
		if index < len(list1Vals) {
			l1Val = list1Vals[index]
		}

		rightChild := new(ListNode)
		rightChild.Val = (l1Val + l2.Next.Val + carry) / 1 % 10
		leftChild.Next = rightChild
		leftChild = rightChild

		if l1Val+l2.Next.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l2 = l2.Next
	}

	if index < len(list1Vals) {
		index++
		for ; index < len(list1Vals); index++ {
			rightChild := new(ListNode)
			rightChild.Val = list1Vals[index]
			leftChild.Next = rightChild
			leftChild = rightChild
		}
	}

	return res
}
