package Easy

import (
	"fmt"
	"testing"
)

func Test70(t *testing.T) {
	fmt.Printf("%v\n", climbStairs(5))
}

func Test746(t *testing.T) {
	fmt.Printf("%v\n", minCostClimbingStairs([]int{1, 100}))
}

func Test70_1(t *testing.T) {
	fmt.Printf("%v\n", climbStairs1(5))
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Printf("%v\n", findLengthOfLCIS1([]int{1, 3, 5, 4, 7}))
}
