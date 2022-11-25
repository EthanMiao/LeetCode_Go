package BackTracking

import "sort"

// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0040.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII.md
var ans1 [][]int
var path1 []int

func combinationSum2(candidates []int, target int) [][]int {
	ans1 = make([][]int, 0)
	path1 = make([]int, 0)
	if len(candidates) == 0 {
		return ans1
	}
	sort.Ints(candidates)
	backTrackingSum2(candidates, target, 0, 0)
	return ans1
}

func backTrackingSum2(candidates []int, target, index, sum int) {
	if sum > target {
		return
	}
	if sum == target {
		temp := make([]int, len(path1))
		copy(temp, path1)
		ans1 = append(ans1, temp)
		return
	}
	for i := index; i < len(candidates); i++ {
		if i-1 >= index && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		path1 = append(path1, candidates[i])
		backTrackingSum2(candidates, target, i+1, sum)
		sum -= candidates[i]
		path1 = path1[:len(path1)-1]
	}
}
