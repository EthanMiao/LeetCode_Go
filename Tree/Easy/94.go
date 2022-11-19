package Easy

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 中序遍历 左中右
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
