/*
#
# @Time : 2019/9/25 16:09
# @Author : Qian Lu
*/

package tree_inorder_traversal

/**
给定一个二叉树，返回它的中序 遍历。

示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res = append(res[0:0])
	RecursionNode(root)
	return res
}

var res []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RecursionNode(node *TreeNode) {
	// left -> root -> right
	if node.Left != nil {
		RecursionNode(node.Left)
	}
	res = append(res, node.Val)
	if node.Right != nil {
		RecursionNode(node.Right)
	}
}
