package t652_duplicate_sub_trees

import (
	"fmt"
	"testing"
)

func Test_t652(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1

	rLeft := new(TreeNode)
	rLeft.Val = 2
	rRight := new(TreeNode)
	rRight.Val = 3
	root.Left = rLeft
	root.Right = rRight

	rlLeft := new(TreeNode)
	rlLeft.Val = 4
	rLeft.Left = rlLeft

	rrLeft := new(TreeNode)
	rrLeft.Val = 2
	rrRight := new(TreeNode)
	rrRight.Val = 4
	rRight.Left = rrLeft
	rRight.Right = rrRight

	rrlLeft := new(TreeNode)
	rrlLeft.Val = 4
	rrLeft.Left = rrlLeft

	res := findDuplicateSubtrees(root)
	fmt.Println(res)
}

func Test_t652_2(t *testing.T) {
	// [0, 0,0 ,0,null, null,0, null,null, 0,0]
	//	       0
	//	   0        0
	//   0   -    -   0
	//  - -          0 0

	//  0,,
	//	0,0,,,
	//	0,,
	//	0,0,,,
	//	0,,0,0,,,
	//	0,0,0,,,,0,,0,0,,,

	root := new(TreeNode)
	root.Val = 0

	rLeft := new(TreeNode)
	rLeft.Val = 0
	rRight := new(TreeNode)
	rRight.Val = 0
	root.Left = rLeft
	root.Right = rRight

	rlLeft := new(TreeNode)
	rlLeft.Val = 0
	rLeft.Left = rlLeft

	rrRight := new(TreeNode)
	rrRight.Val = 0
	rRight.Right = rrRight

	rrlLeft := new(TreeNode)
	rrlLeft.Val = 0
	rrlRight := new(TreeNode)
	rrlRight.Val = 0
	rrRight.Left = rrlLeft
	rrRight.Left = rrlRight

	res := findDuplicateSubtrees(root)
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
