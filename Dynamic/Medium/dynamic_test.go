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
