/*
#
# @Time : 2019/9/29 15:55
# @Author : Qian Lu
*/

package tree_balance

import (
	"fmt"
	"testing"
)

func Test_TreeIsBalance(t *testing.T) {
	/**
	示例 1:
	给定二叉树 [3,9,20,null,null,15,7]
		3
	   / \
	  9  20
		/  \
	   15   7
	返回 true
	*/

	root := new(TreeNode)
	root.Val = 3

	rLeft := new(TreeNode)
	rLeft.Val = 9
	rRight := new(TreeNode)
	rRight.Val = 20
	root.Left = rLeft
	root.Right = rRight

	rrLeft := new(TreeNode)
	rrLeft.Val = 15
	rrRight := new(TreeNode)
	rrRight.Val = 7
	rRight.Left = rrLeft
	rRight.Right = rrRight

	fmt.Println(isBalanced(root))

	/**
	示例 2:
	给定二叉树 [1,2,2,3,3,null,null,4,4]
		   1
		  / \
		 2   2
		/ \
	   3   3
	  / \
	 4   4
	返回 false 。
	*/

	root2 := new(TreeNode)
	root2.Val = 1

	r2Left := new(TreeNode)
	r2Left.Val = 2
	r2Right := new(TreeNode)
	r2Right.Val = 2
	root2.Left = r2Left
	root2.Right = r2Right

	r2lLeft := new(TreeNode)
	r2lLeft.Val = 3
	r2lRight := new(TreeNode)
	r2lRight.Val = 3
	r2Left.Left = r2lLeft
	r2Left.Right = r2lRight

	r2llLeft := new(TreeNode)
	r2llLeft.Val = 4
	r2llRight := new(TreeNode)
	r2llRight.Val = 4
	r2lLeft.Left = r2llLeft
	r2lLeft.Right = r2llRight

	fmt.Println(isBalanced(root2))
}
