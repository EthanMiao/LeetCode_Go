package Medium

// https://leetcode.cn/problems/maximum-binary-tree/description/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	for i, val := range nums {
		if val > nums[index] {
			index = i
		}
	}
	node := &TreeNode{Val: nums[index]}
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return node
}
