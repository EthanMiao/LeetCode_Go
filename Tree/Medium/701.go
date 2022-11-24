package Medium

// https://leetcode.cn/problems/insert-into-a-binary-search-tree/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	for node != nil {
		if node.Val < val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				break
			}
			node = node.Right
		}
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				break
			}
			node = node.Left
		}
	}
	return root
}
