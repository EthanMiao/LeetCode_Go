package Easy

import "LeetCode_Go/Tree/Medium"

// https://leetcode.cn/problems/merge-two-binary-trees/
func mergeTrees(root1 *Medium.TreeNode, root2 *Medium.TreeNode) *Medium.TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
