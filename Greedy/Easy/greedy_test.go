package Easy

import (
	"fmt"
	"testing"
)

func Test455(t *testing.T) {
	fmt.Printf("%v\n", findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}

func Test1005(t *testing.T) {
	fmt.Printf("%v\n", largestSumAfterKNegations([]int{4, 2, 3}, 1))
}

func Test860(t *testing.T) {
	fmt.Printf("%v\n", lemonadeChange([]int{5, 5, 10, 10, 20}))
}
