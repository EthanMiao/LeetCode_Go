package Medium

import "fmt"

// https://leetcode.cn/problems/unique-binary-search-trees/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum = sum + dp[j-1]*dp[i-j]
		}
		dp[i] = sum
		fmt.Printf("%v\n", dp)
	}
	return dp[n]
}
