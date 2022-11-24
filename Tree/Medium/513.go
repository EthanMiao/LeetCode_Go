package Medium

// https://leetcode.cn/problems/find-bottom-left-tree-value/
var curHeight = 0
var result = 0

// DFS
func findBottomLeftValue(root *TreeNode) int {
	depthFirst(root, 0)
	return result
}

func depthFirst(node *TreeNode, height int) {
	if node == nil {
		return
	}
	height++
	depthFirst(node.Left, height)
	depthFirst(node.Right, height)
	if curHeight < height {
		curHeight = height
		result = node.Val
	}
	return
}

// BFS
func findBottomLeftValue1(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		ans = node.Val
	}
	return ans
}
