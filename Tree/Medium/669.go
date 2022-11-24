package Medium

// https://leetcode.cn/problems/trim-a-binary-search-tree/
// 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > high:
		return trimBST(root.Left, low, high)
	case root.Val < low:
		return trimBST(root.Right, low, high)
	default:
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
}

// 迭代
func trimBST1(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for node := root; node.Left != nil; {
		if node.Left.Val < low {
			node.Left = node.Left.Right
		} else {
			node = node.Left
		}
	}
	for node := root; node.Right != nil; {
		if node.Right.Val > high {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}
