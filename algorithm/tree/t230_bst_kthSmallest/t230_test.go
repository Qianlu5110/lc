package t230_bst_kthSmallest

import (
	"fmt"
	"testing"
)

func Test_t230(t *testing.T) {
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

	res := kthSmallest(root, 2)
	fmt.Println(res)
}
