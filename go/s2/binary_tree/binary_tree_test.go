package binarytree

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	bf := BreadthFirstTree(4, 2, 7, 1, 3)
	PrintTree(bf)

	// b := T2N(bf)
	// PrintBN(b)
	// fmt.Println("x2")
}
