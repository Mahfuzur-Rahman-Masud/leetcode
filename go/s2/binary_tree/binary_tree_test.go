package binarytree

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	b := IsSameTree(BreadthFirstTree(4, 2, 7, 1, 3), Make(4, 2, 7, 1, 3))
	fmt.Println(b)
}
