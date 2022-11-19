package Easy

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 前序遍历 中左右

// 递归
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// 遍历
func preorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	st := list.New()
	st.PushFront(root)

	for st.Len() > 0 {
		node := st.Remove(st.Front()).(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			st.PushFront(node.Right)
		}
		if node.Left != nil {
			st.PushFront(node.Left)
		}
	}
	return result
}
