package t230_bst_kthSmallest

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第k个最小元素（从 1 开始计数）

示例 1：
    3
   / \
  1   4
   \
    2
输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：

        5
       / \
      3   6
     / \
    2   4
   /
  1
输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3


提示：

树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104


进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	min := -1
	minIndex := k

	var smallest func(node *TreeNode)
	smallest = func(node *TreeNode) {
		if node == nil {
			return
		}
		smallest(node.Left)
		if minIndex == 0 {
			return
		}

		// 递归中只处理当前节点就可以了，有叶子节点 就进一步递归
		if node.Val >= min {
			min = node.Val
			minIndex--
		}
		smallest(node.Right)
	}

	smallest(root)
	return min
}

/**
  这个算法 体现了自身在树递归环节的缺陷
*/
func kthSmallestOld(root *TreeNode, k int) int {
	min := -1
	minIndex := 0

	flag := false
	var smallest func(node *TreeNode) int
	smallest = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		temp := smallest(node.Left)
		if flag {
			return temp
		}
		if node.Val < temp && temp >= 0 && !flag {
			min = temp //3
			minIndex++
			if minIndex == k {
				flag = true
				return min
			}
		}

		min = node.Val
		minIndex++
		if minIndex == k {
			flag = true
			return min
		}

		templ := smallest(node.Right)
		if flag {
			return templ
		}
		if min > templ && templ >= 0 && !flag {
			min = templ // 2
			minIndex++
			if minIndex == k {
				flag = true
				return min
			}
		}

		return min
	}

	min = smallest(root)

	return min
}
