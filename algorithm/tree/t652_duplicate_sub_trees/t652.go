package t652_duplicate_sub_trees

import (
	"fmt"
	"strconv"
)

/**
652. 寻找重复的子树
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：
1_2_4_4_#_#_#_#_3_2_#_#_4_#_#

        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
下面是两个重复的子树：

      2
     /
    4
和
    4
因此，你需要以列表的形式返回上述重复子树的根结点。

*/

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	var result []*TreeNode

	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		serial := fmt.Sprintf("%d,%s,%s", root.Val, dfs(root.Left), dfs(root.Right))
		fmt.Println(serial)

		m[serial]++
		if m[serial] == 2 {
			result = append(result, root)
		}
		return serial
	}
	dfs(root)
	return result

	//nodes = []*TreeNode{}
	//findDuplicateSubtree(root)
	//return nodes
}

var tempMap = make(map[string]int)
var nodes []*TreeNode

func findDuplicateSubtree(node *TreeNode) string {
	if node == nil {
		return "#"
	}
	var leftStr, rightStr string
	leftStr = findDuplicateSubtree(node.Left)
	rightStr = findDuplicateSubtree(node.Right)

	nodeStr := "[" + strconv.Itoa(node.Val) + "_" + leftStr + "_" + rightStr + "]"
	tempMap[nodeStr]++
	fmt.Println(nodeStr)
	if tempMap[nodeStr] == 2 {
		nodes = append(nodes, node)
	}
	return nodeStr
}
