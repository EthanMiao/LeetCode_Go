package Medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/house-robber-iii/
func rob2(root *TreeNode) int {
	result := robNode(root)
	return max(result[0], result[1])
}

func robNode(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	left, right := robNode(node.Left), robNode(node.Right)
	selected := max(left[0], left[1]) + max(right[0], right[1])
	notSelected := node.Val + left[0] + right[0]
	return []int{selected, notSelected}
}
