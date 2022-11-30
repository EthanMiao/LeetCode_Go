package Medium

import (
	"fmt"
	"testing"
)

func Test376(t *testing.T) {
	fmt.Printf("%v\n", wiggleMaxLength([]int{1, 2, 3}))
}

func Test53(t *testing.T) {
	fmt.Printf("%v\n", maxSubArray([]int{5, 4, -1, 7, 8}))
}

func Test55(t *testing.T) {
	fmt.Printf("%v\n", jump([]int{2, 3, 1, 1, 2}))
}

func Test134(t *testing.T) {
	fmt.Printf("%v\n", canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))
}

func Test406(t *testing.T) {
	fmt.Printf("%v\n", reconstructQueue([][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}))
}

func Test452(t *testing.T) {
	fmt.Printf("%v\n", findMinArrowShots([][]int{{1, 2}, {4, 5}, {1, 5}}))
}

func Test435(t *testing.T) {
	fmt.Printf("%v\n", findMinArrowShots([][]int{{1, 2}, {2, 3}, {1, 3}}))
}

func Test56(t *testing.T) {
	fmt.Printf("%v\n", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Printf("%v\n", merge([][]int{{1, 4}, {2, 3}}))
}

func Test738(t *testing.T) {
	fmt.Printf("%v\n", monotoneIncreasingDigits(10))
}
