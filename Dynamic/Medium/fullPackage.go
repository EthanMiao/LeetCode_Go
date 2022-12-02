package Medium

// 先遍历物品，再遍历背包
func fullPackage(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

// 先遍历背包，再遍历物品
func fullPackage1(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for i := 0; i <= bagWeight; i++ {
		for j := 0; j < len(weight); j++ {
			if i >= weight[j] {
				dp[i] = max(dp[i], dp[i-weight[j]]+value[j])
			}
		}
	}
	return dp[bagWeight]
}
