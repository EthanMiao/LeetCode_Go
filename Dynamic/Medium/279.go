package Medium

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/perfect-squares/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	root := int(math.Sqrt(float64(n)))
	for i := 1; i <= root; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j-i*i] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
		fmt.Printf("%v\n", dp)
	}
	if dp[n] == math.MaxInt {
		return -1
	}
	return dp[n]
}
