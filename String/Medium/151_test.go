package Medium

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	result := reverseWords(" hello test world ")
	fmt.Println(result)
}

func TestReverseWords2(t *testing.T) {
	result := reverseWords(" hello world ")
	fmt.Println("$" + result + "$")
}
