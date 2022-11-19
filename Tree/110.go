package Tree

// 自顶向下
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(depth(root.Left), depth(root.Right)) <= 1 && isBalanced1(root.Left) && isBalanced1(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 自下而上
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight, rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
