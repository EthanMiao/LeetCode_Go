package BackTracking

// https://leetcode.cn/problems/combination-sum-iii/description/
var ansSum [][]int

func combinationSum3(k int, n int) [][]int {
	ansSum = make([][]int, 0)
	backTrackingSum(k, n, 1, []int{})
	return ansSum
}

func backTrackingSum(k, n, start int, track []int) {
	if len(track) == k {
		if sum(track) == n {
			temp := make([]int, k)
			copy(temp, track)
			ansSum = append(ansSum, temp)
		}
		return
	}
	for i := start; i <= 9-(k-len(track))+1; i++ {
		track = append(track, i)
		backTrackingSum(k, n, i+1, track)
		track = track[:len(track)-1]
	}
}

func sum(track []int) int {
	ans := 0
	for _, v := range track {
		ans += v
	}
	return ans
}
