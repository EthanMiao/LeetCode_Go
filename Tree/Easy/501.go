package Easy

// https://leetcode.cn/problems/find-mode-in-binary-search-tree/
func findMode(root *TreeNode) []int {
	answer := make([]int, 0)
	curr, count, maxCount := 0, 0, 0

	var update func(int)
	update = func(val int) {
		if val == curr {
			count++
		} else {
			curr, count = val, 1
		}
		if count == maxCount {
			answer = append(answer, val)
		}
		if count > maxCount {
			maxCount = count
			answer = []int{curr}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return answer
}
