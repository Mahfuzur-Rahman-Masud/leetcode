package binarytree

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	t1 := BreadthFirstTree(1, 3, 2, 5)
	t2 := BreadthFirstTree(2, 1, 3, -1, 4, -1, 7)

	PrintTree(t1)
	fmt.Println()
	fmt.Println()
	PrintTree(t2)

	fmt.Println()
	fmt.Println()
	t3 := mergeTrees(t1, t2)
	PrintTree(t3)

}
