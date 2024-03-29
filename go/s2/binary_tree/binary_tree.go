package binarytree

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BreadthFirstTree(val ...int) *TreeNode {
	if val == nil {
		return nil
	}

	t := &TreeNode{
		Val: val[0],
	}

	slice := []*TreeNode{t}
	mark := 0
	left := true

	for i, v := range val {
		if i == 0 {
			continue
		}

		tree := slice[mark]

		if left && tree.Left == nil && v != -1 {
			tree.Left = &TreeNode{Val: v}
			left = false
			slice = append(slice, tree.Left)
		} else if left && v == -1 {
			left = false
		} else if !left && v != -1 && tree.Right == nil {
			tree.Right = &TreeNode{Val: v}
			left = false
			slice = append(slice, tree.Right)
			mark++
			left = true
		} else if !left {
			mark++
			left = true
		}
	}

	return t
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right))
}

func PrintTree(root *TreeNode) {
	x := toArray(root)

	for _, x := range x {
		for _, x := range x {
			fmt.Printf("   %s   ", x)
		}
		fmt.Println()
		fmt.Println()
	}
}

func toArray(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}

	h := getHeight(root)
	col := int(math.Pow(2, float64(h+1)) - 1)
	res := make([][]string, h+1)

	for i := 0; i < h+1; i++ {
		row := make([]string, col)
		// init res 2d arr
		for j := 0; j < col; j++ {
			row[j] = ""
		}
		res[i] = row
	}

	fillNode(root, 0, col, 0, res)

	return res
}

func fillNode(n *TreeNode, l, r, h int, res [][]string) {
	if n == nil {
		return
	}

	var mid int = (l + r) / 2
	res[h][mid] = strconv.Itoa(n.Val)

	if n.Left != nil {
		fillNode(n.Left, l, mid, h+1, res)
	}

	if n.Right != nil {
		fillNode(n.Right, mid+1, r, h+1, res)
	}
}

func getHeight(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	var lH, rH int
	if root.Left != nil {
		lH = getHeight(root.Left) + 1
	}

	if root.Right != nil {
		rH = getHeight(root.Right) + 1
	}

	if lH > rH {
		return lH
	}

	return rH
}
