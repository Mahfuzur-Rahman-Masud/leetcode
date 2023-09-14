package leet

import (
	"fmt"
	"testing"
)

func TestPreorder(t *testing.T) {
	r := &Node{Val: 1, Children: []*Node{
		{Val: 2, Children: []*Node{{Val: 4}}}, {Val: 3},
	}}

	// fmt.Println(r)

	out := preorder(r)
	fmt.Println("result:", out)
}

func printNaryTree(t *Node) {
	if t == nil {
		return
	}

	fmt.Println(t.Val)

	for _, c := range t.Children {
		printNaryTree(c)
	}

}
