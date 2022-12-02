package Easy

// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// 完全背包解法
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	value := []int{1, 2}
	for i := 0; i <= n; i++ {
		for j := 0; j < 2; j++ {
			if i >= value[j] {
				dp[i] += dp[i-value[j]]
			}
		}
	}
	return dp[n]
}
