package String

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	result := ReverseWords(" hello test world ")
	fmt.Println(result)
}

func TestReverseWords2(t *testing.T) {
	result := ReverseWords(" hello world ")
	fmt.Println("$" + result + "$")
}
