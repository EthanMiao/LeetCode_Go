package BackTracking

// https://leetcode.cn/problems/combinations/
var ans [][]int
var track []int

func combine(n int, k int) [][]int {
	ans = make([][]int, 0)
	track = make([]int, 0)
	if n <= 0 || k <= 0 || k > n {
		return ans
	}
	backTracking(n, k, 1)
	return ans
}

func backTracking(n, k, start int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		ans = append(ans, temp)
		return
	}
	if len(track)+n-start+1 < k {
		return
	}
	for i := start; i <= n-(k-len(track))+1; i++ {
		track = append(track, i)
		backTracking(n, k, i+1)
		track = track[:len(track)-1]
	}
}
