package t116_tree_connect_right_point

import (
	"fmt"
	"testing"
)

func Test_findRightPoint(t *testing.T) {
	/**
		示例 1:
		给定一个 完美二叉树 [1,2,3,4,5,6,7]
		    1
		   /  \
		  2    3
		 / \  / \
		4  5  6  7
	 	返回 [1,#,2,3,#,4,5,6,7,#]
	*/
	root := new(Node)
	root.Val = 1

	rLeft := new(Node)
	rLeft.Val = 2
	rRight := new(Node)
	rRight.Val = 3
	root.Left = rLeft
	root.Right = rRight

	rlLeft := new(Node)
	rlLeft.Val = 4
	rlRight := new(Node)
	rlRight.Val = 5
	rLeft.Left = rlLeft
	rLeft.Right = rlRight

	rrLeft := new(Node)
	rrLeft.Val = 6
	rrRight := new(Node)
	rrRight.Val = 7
	rRight.Left = rrLeft
	rRight.Right = rrRight

	res := connect(root)
	fmt.Println(res.Val)
}
