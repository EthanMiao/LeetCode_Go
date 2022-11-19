package Easy

// https://leetcode.cn/problems/binary-tree-postorder-traversal/
// 后序遍历 左右中

// 递归
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// 遍历
func postorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	// undo

	return result
}
