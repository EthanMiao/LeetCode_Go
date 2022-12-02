package Medium

import "fmt"

// https://leetcode.cn/problems/ones-and-zeroes/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeroNum, oneNum := 0, 0
		for _, v := range str {
			if v == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}

		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
		//print
		fmt.Println("#########")
		for _, d := range dp {
			fmt.Printf("%v\n", d)
		}
		fmt.Println()
	}
	return dp[m][n]
}
