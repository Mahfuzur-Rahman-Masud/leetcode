package binarytree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSecondMinimumValue(t *testing.T) {
	bt := BreadthFirstTree(2, 2, 5, -1, -1, 5, 7)
	PrintTree(bt)
	r := findSecondMinimumValue(bt)
	assert.Equal(t, 5, r)

	assert.Equal(t, -1, findSecondMinimumValue(BreadthFirstTree(2, 2, 2)))
}

func findSecondMinimumValue(root *TreeNode) int {
	res := &[2]int{math.MaxInt, math.MaxInt}
	findSecondMinimumValue_res(root, res)
	if res[1] == math.MaxInt {
		return -1
	}
	return res[1]
}

func findSecondMinimumValue_res(root *TreeNode, res *[2]int) {

	if root == nil {
		return
	}

	if root.Val < res[0] {
		res[0], res[1] = root.Val, res[0]
	} else if root.Val < res[1] && root.Val > res[0] {
		res[1] = root.Val
	}

	findSecondMinimumValue_res(root.Left, res)
	findSecondMinimumValue_res(root.Right, res)
}
