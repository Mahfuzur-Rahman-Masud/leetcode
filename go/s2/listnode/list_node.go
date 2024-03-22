package listnode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) ToSlice() []int {
	out := []int{}
	if n == nil {
		return out
	}

	node := n
	for node != nil {
		out = append(out, node.Val)
		node = node.Next
	}

	return out
}

func (n *ListNode) ToString() string {
	if n == nil {
		return "[]"
	}
	node := n
	out := []string{}
	for node != nil {
		out = append(out, strconv.Itoa(node.Val))
		node = node.Next
	}

	return "[" + strings.Join(out, ", ") + "]"
}

func ListOf(data ...int) *ListNode {
	return CreateLLFrom(data)
}

func SliceOf(data ...int) []int {
	return data
}

func CreateLLFrom(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}

	root := &ListNode{
		Val: data[0],
	}

	last := root
	for i, v := range data {
		if i == 0 {
			continue
		}

		ln := ListNode{
			Val: v,
		}

		last.Next = &ln
		last = &ln
	}

	return root
}
