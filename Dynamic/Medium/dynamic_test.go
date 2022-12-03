package Medium

import (
	"fmt"
	"testing"
)

func Test62(t *testing.T) {
	fmt.Printf("%v\n", uniquePaths(3, 7))
}

func Test63(t *testing.T) {
	fmt.Printf("%v\n", uniquePathsWithObstacles([][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func Test343(t *testing.T) {
	fmt.Printf("%v\n", integerBreak(10))
}

func Test96(t *testing.T) {
	fmt.Printf("%v\n", numTrees(3))
}

func TestDemo01Package(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Printf("%v\n", Demo01Package(weight, value, 4))
}

func TestDemo01Package1(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Printf("%v\n", Demo01Package1(weight, value, 4))
}

func TestCanPartition(t *testing.T) {
	fmt.Printf("%v\n", canPartition([]int{1, 5, 11, 5}))
}

func TestLastStoneWeightII(t *testing.T) {
	fmt.Printf("%v\n", lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Printf("%v\n", lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Printf("%v\n", findTargetSumWays([]int{2, 3, 1, 2, 1}, 3))
}

func TestFindMaxForm(t *testing.T) {
	fmt.Printf("%v\n", findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

func TestChange(t *testing.T) {
	fmt.Printf("%v\n", change(5, []int{1, 2, 5}))
}

func TestCombinationSum4(t *testing.T) {
	fmt.Printf("%v\n", combinationSum4([]int{1, 2, 3}, 4))
}

func TestCoinChange(t *testing.T) {
	fmt.Printf("%v\n", coinChange([]int{1, 2, 5}, 5))
}

func TestNumSquares(t *testing.T) {
	fmt.Printf("%v\n", numSquares(12))
}

func TestWordBreak(t *testing.T) {
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Printf("%v\n", wordBreak("applepenapple", []string{"apple", "pen"}))
}

func TestRob(t *testing.T) {
	fmt.Printf("%v\n", rob([]int{2, 7, 9, 3, 1}))
}

func TestRob1(t *testing.T) {
	fmt.Printf("%v\n", rob1([]int{2, 3, 2}))
}

func TestMaxSubArray(t *testing.T) {
	fmt.Printf("%v\n", maxSubArray([]int{-1, -2}))
}
