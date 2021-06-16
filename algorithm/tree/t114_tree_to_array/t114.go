package t114_tree_to_array

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	linkLeftDown := root.Left
	if linkLeftDown != nil {
		for linkLeftDown.Right != nil {
			linkLeftDown = linkLeftDown.Right
		}
		linkLeftDown.Right = root.Right
		root.Right = root.Left
	}
	root.Left = nil
}
