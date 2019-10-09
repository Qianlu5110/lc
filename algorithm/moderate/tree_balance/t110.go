/*
#
# @Time : 2019/9/29 15:55
# @Author : Qian Lu
*/

package tree_balance

import (
	"fmt"
	"math"
)

/**
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。

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

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func isBalanced(root *TreeNode) bool {
	return NodeBalanced(root) != -1
}

func NodeBalanced(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	left := NodeBalanced(node.Left)
	right := NodeBalanced(node.Right)
	fmt.Println(fmt.Sprintf("val:%d，left:%f ,right:%f", node.Val, left, right))
	if left == -1 || right == -1 {
		return -1
	}

	fmt.Println(fmt.Sprintf("val:%d，left:%f, right:%f, abs:%f", node.Val, left, right, math.Abs(left-right)))
	if math.Abs(left-right) <= 1 {
		return math.Max(left, right) + 1
	} else {
		return -1
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
