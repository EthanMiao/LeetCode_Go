package Medium

import "math"

// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 递归
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
}

// 迭代 中序遍历
func isValidBST1(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	inorder := math.MinInt
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
