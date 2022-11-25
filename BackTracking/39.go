package BackTracking

// https://leetcode.cn/problems/combination-sum/
var paths [][]int
var path []int

func combinationSum(candidates []int, target int) [][]int {
	paths = make([][]int, 0)
	if len(candidates) == 0 {
		return paths
	}
	backTrackSum(candidates, target, 0, 0)
	return paths
}

func backTrackSum(candidates []int, target, index, sum int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		paths = append(paths, temp)
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		backTrackSum(candidates, target, i, sum)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
