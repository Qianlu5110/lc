package t114_tree_to_array

import (
	"fmt"
	"testing"
)

func Test_t144(t *testing.T) {

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

	//rrLeft := new(TreeNode)
	//rrLeft.Val = 15
	rrRight := new(TreeNode)
	rrRight.Val = 7
	//rRight.Left = rrLeft
	rRight.Right = rrRight

	flatten(root)
	fmt.Print("over")

}
