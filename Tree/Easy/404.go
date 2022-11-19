package Easy

// https://leetcode.cn/problems/sum-of-left-leaves/
// DFS
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

func dfs(node *TreeNode) int {
	ans := 0
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return ans
}

func isLeafNode(node *TreeNode) bool {
	return node.Right == nil && node.Left == nil
}

// BFS
func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			if isLeafNode(node.Left) {
				ans += node.Left.Val
			} else {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			queue = append(queue, node.Right)
		}
	}
	return ans
}
