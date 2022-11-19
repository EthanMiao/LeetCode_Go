package Tree

import "strconv"

// 官方，省内存
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = make([]string, 0)
	constructPaths(root, "")
	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, path)
		} else {
			path += "->"
			constructPaths(root.Left, path)
			constructPaths(root.Right, path)
		}
	}
}

// 自有
func binaryTreePaths1(root *TreeNode) []string {
	return buildPath(root, "")
}

func buildPath(root *TreeNode, path string) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}

	if path == "" {
		path = strconv.FormatInt(int64(root.Val), 10)
	} else {
		path = path + "->" + strconv.FormatInt(int64(root.Val), 10)
	}

	if root.Left == nil && root.Right == nil {
		return append(result, path)
	}

	result = append(result, buildPath(root.Left, path)...)
	result = append(result, buildPath(root.Right, path)...)
	return result
}
