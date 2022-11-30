package Easy

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	for i := 0; i < k; i++ {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 优化
func largestSumAfterKNegations1(nums []int, k int) int {
	sort.Slice(nums, func(a, b int) bool {
		return math.Abs(float64(nums[a])) > math.Abs(float64(nums[b]))
	})
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
