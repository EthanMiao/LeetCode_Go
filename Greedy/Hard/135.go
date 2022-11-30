package Hard

import "fmt"

// https://leetcode.cn/problems/candy/description/
func candy(ratings []int) int {
	n := len(ratings)
	forward := make([]int, n)
	forward[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			forward[i] = forward[i-1] + 1
		} else {
			forward[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			forward[i] = max(forward[i], forward[i+1]+1)
		}
	}
	sum := 0
	fmt.Printf("%v\n", forward)
	for _, v := range forward {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
