package t538_tree_bst_convert_to_greater

import (
	"fmt"
	"testing"
)

/**
示例 1：
    3
   / \
  1   4
   \
    2

示例 1：
    7
   / \
  10   4
   \
    9

*/
func Test_538(t *testing.T) {
	root := new(TreeNode)
	root.Val = 3

	rLeft := new(TreeNode)
	rLeft.Val = 1
	rRight := new(TreeNode)
	rRight.Val = 4
	root.Left = rLeft
	root.Right = rRight

	rlLeft := new(TreeNode)
	rlLeft.Val = 2
	rLeft.Right = rlLeft

	res := convertBST(root)
	fmt.Println(res)

}
