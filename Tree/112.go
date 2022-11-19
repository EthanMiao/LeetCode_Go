package Tree

// DFS 第一种写法
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// DFS 第二种写法
func hasPathSum1(root *TreeNode, targetSum int) bool {
	return dfsSum(root, targetSum, 0)
}

func dfsSum(node *TreeNode, targetSum int, sum int) bool {
	if node == nil {
		return false
	}
	sum += node.Val
	if node.Left == nil && node.Right == nil && sum == targetSum {
		return true
	}
	return dfsSum(node.Left, targetSum, sum) || dfsSum(node.Right, targetSum, sum)
}

// BFS
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	qNode := make([]*TreeNode, 0)
	qNode = append(qNode, root)
	qSum := make([]int, 0)
	qSum = append(qSum, root.Val)
	for len(qNode) > 0 {
		node := qNode[0]
		qNode = qNode[1:]
		sum := qSum[0]
		qSum = qSum[1:]
		if node.Left == nil && node.Right == nil {
			// 剪枝
			if sum == targetSum {
				return true
			}
			continue
		}
		if node.Left != nil {
			qNode = append(qNode, node.Left)
			qSum = append(qSum, sum+node.Left.Val)
		}
		if node.Right != nil {
			qNode = append(qNode, node.Right)
			qSum = append(qSum, sum+node.Right.Val)
		}
	}
	return false
}
