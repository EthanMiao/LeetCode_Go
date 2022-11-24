package Medium

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	ancestor := root
	for ancestor != nil {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			break
		}
	}
	return ancestor
}
