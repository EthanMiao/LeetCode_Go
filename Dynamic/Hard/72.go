package Hard

import "math"

// https://leetcode.cn/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minList([]int{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]}) + 1
			}
		}
	}
	return dp[m][n]
}

func minList(nums []int) int {
	result := math.MaxInt
	for _, num := range nums {
		if num < result {
			result = num
		}
	}
	return result
}
