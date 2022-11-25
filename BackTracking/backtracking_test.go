package BackTracking

import (
	"fmt"
	"testing"
)

func Test131(t *testing.T) {
	fmt.Printf("%v\n", partition("bb"))
}

func Test93(t *testing.T) {
	fmt.Printf("%v\n", restoreIpAddresses("101023"))
	fmt.Printf("%v\n", restoreIpAddresses("0000"))
	fmt.Printf("%v\n", restoreIpAddresses("25525511135"))
}

func Test78(t *testing.T) {
	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))
}

func Test77(t *testing.T) {
	combine(4, 2)
}

func Test90(t *testing.T) {
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 2, 2}))
	fmt.Printf("%v\n", subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func Test491(t *testing.T) {
	fmt.Printf("%v\n", findSubsequences([]int{4, 6, 7, 7}))
	fmt.Printf("%v\n", findSubsequences([]int{4, 4, 3, 2, 1}))
	fmt.Printf("%v\n", findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 1, 1}))
}

func Test46(t *testing.T) {
	fmt.Printf("%v\n", permute([]int{1, 2, 3}))
}

func Test47(t *testing.T) {
	fmt.Printf("%v\n", permuteUnique([]int{3, 3, 0, 3}))
}

func Test216(t *testing.T) {
	fmt.Printf("%v\n", combinationSum3(3, 9))
}
