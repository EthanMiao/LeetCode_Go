package Medium

import "fmt"

func Demo01Package(weight, value []int, bagWeight int) int {
	length := len(weight)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, bagWeight+1)
	}

	// 初始化
	for i := weight[0]; i <= bagWeight; i++ {
		dp[0][i] = value[0]
	}

	// 逐行打印数组
	fmt.Println("##################")
	for i := 0; i < length; i++ {
		fmt.Printf("%v\n", dp[i])
	}
	fmt.Println()

	for i := 1; i < length; i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}

		// 逐行打印数组
		fmt.Println("##################")
		for i := 0; i < length; i++ {
			fmt.Printf("%v\n", dp[i])
		}
		fmt.Println()
	}

	return dp[length-1][bagWeight]
}

func Demo01Package1(weight, value []int, bagWeight int) int {
	length := len(weight)
	dp := make([]int, bagWeight+1)

	for j := weight[0]; j <= bagWeight; j++ {
		dp[j] = value[0]
	}
	fmt.Printf("%v\n", dp)

	for i := 1; i < length; i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
		fmt.Printf("%v\n", dp)
	}

	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
