package Medium

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := &TreeNode{}

	var dfs func(root, p, q *TreeNode) bool
	dfs = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}
		lson := dfs(root.Left, p, q)
		rson := dfs(root.Right, p, q)
		if (lson && rson) || ((root.Val == p.Val || root.Val == q.Val) && (lson || rson)) {
			ans = root
		}
		return lson || rson || root.Val == p.Val || root.Val == q.Val
	}

	dfs(root, p, q)
	return ans
}
