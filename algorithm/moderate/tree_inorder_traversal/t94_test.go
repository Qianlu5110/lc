/*
#
# @Time : 2019/9/25 16:09
# @Author : Qian Lu
*/

package tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test_InorderTraversal(t *testing.T) {
	/**
	输入: [1,null,2,3]
	   1
		\
		 2
		/
	   3
	输出: [1,3,2]
	*/

	root := new(TreeNode)
	root.Val = 1
	rcRight := new(TreeNode)
	rcRight.Val = 2
	root.Right = rcRight
	rcrLeft := new(TreeNode)
	rcrLeft.Val = 3
	rcRight.Left = rcrLeft

	fmt.Println(InorderTraversal(root))

	/**
	输入: [1]
	   1
	输出: [1]
	*/

	r := new(TreeNode)
	r.Val = 1
	fmt.Println(InorderTraversal(r))
}
