package Easy

// 二叉搜索树满足如下性质
// 左子树所有节点的元素值均小于根的元素值
// 右子树所有节点的元素值均大于根的元素值

// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
