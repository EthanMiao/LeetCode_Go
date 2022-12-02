package Medium

import "fmt"

// https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	wordDictMap := map[string]bool{}
	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	length := len(s)
	dp := make([]bool, length+1)
	dp[0] = true

	// 遍历背包
	for i := 1; i <= len(s); i++ {
		// 遍历单词
		for j := 0; j < i; j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
		fmt.Printf("%v\n", dp)
	}
	return dp[length]
}
