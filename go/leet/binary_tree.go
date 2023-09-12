package leet

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return tiltTree(root) + findTilt(root.Left) + findTilt(root.Right)
}

func tiltTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return int(math.Abs(float64(sumTree(node.Right)) - float64(sumTree(node.Left))))
}

func sumTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return node.Val + sumTree(node.Left) + sumTree(node.Right)
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, v := maxPathBinaryTree(root, 0)
	return v
}

func maxPathBinaryTree(root *TreeNode, seed int) (int, int) {
	if root == nil {
		return seed, 0
	}

	seed++

	l, vl := maxPathBinaryTree(root.Left, 0)
	r, vr := maxPathBinaryTree(root.Right, 0)

	return mxx(seed+l, seed+r), mxx(l+r, vl, vr)
}

func mxx(val ...int) int {
	max := 0
	for _, v := range val {
		if v > max {
			max = v
		}
	}

	return max
}

func findMode2(root *TreeNode) []int {
	m := make(map[int]int)

	buffer := make([]*TreeNode, 0, 64)
	mark := 0

	buffer = append(buffer, root)

	out := []int{}
	max := 0

	for mark >= 0 {
		n := buffer[mark]
		k := n.Val
		m[k]++
		v := m[k]

		if v == max {
			out = append(out, k)
		} else if v > max {
			max = v
			out = []int{k}
		}

		if n.Left != nil && n.Right != nil {
			buffer[mark] = n.Left
			mark++
			if mark < len(buffer) {
				buffer[mark] = n.Right
			} else {
				buffer = append(buffer, n.Right)
			}
		} else if n.Left != nil {
			buffer[mark] = n.Left
		} else if n.Right != nil {
			buffer[mark] = n.Right
		} else {
			mark--
		}
	}

	return out
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	tallyNodes(root, m)

	out := []int{}
	max := 0

	for k, v := range m {
		if v == max {
			out = append(out, k)
		}

		if v > max {
			max = v
			out = []int{k}
		}
	}

	return out
}

func tallyNodes(r *TreeNode, m map[int]int) {
	if r == nil {
		return
	}

	m[r.Val]++
	tallyNodes(r.Left, m)
	tallyNodes(r.Right, m)
}

func createTreeBreadthFirstTLRArray(val []int) *TreeNode {
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

func sumOfLeftLeaves2(root *TreeNode) int {
	left := []*TreeNode{}
	right := []*TreeNode{root}

	sum := 0
	node := root

	for len(left) != 0 || len(right) != 0 {
		if len(left) != 0 {
			node = left[0]
			left = left[1:]

			if node.Left == nil && node.Right == nil {
				sum += node.Val
			}
		} else {
			node = right[0]
			right = right[1:]
		}

		if node.Left != nil {
			left = append(left, node.Left)
		}

		if node.Right != nil {
			right = append(right, node.Right)
		}
	}

	return sum
}

func sumOfLeftLeaves(root *TreeNode) int {

	return sumLeft(root, false, 0)
}

func sumLeft(root *TreeNode, isLeft bool, sum int) int {
	if root == nil {
		return sum
	}

	if isLeft && root.Left == nil && root.Right == nil {
		return sum + root.Val
	} else if root.Left == nil && root.Right == nil {
		return sum
	}

	return +sumLeft(root.Right, false, sumLeft(root.Left, true, sum))
}

func InorderTraversal2(root *TreeNode) []int {
	out := &[]int{}
	parent := map[*TreeNode]*TreeNode{}
	seen := map[*TreeNode]bool{}

	n := root
	l := root
	// r := root
	node := root

	for n != nil {
		l = n
		for l != nil {
			node = l
			l = l.Left
			parent[l] = node
		}

		*out = append(*out, node.Val)
		seen[node] = true

		if node.Right != nil {
			parent[node.Right] = node
			n = node.Right
			continue
		}

		p, ok := parent[node]
		for ok && p != nil && seen[p] {
			p, ok = parent[p]
		}

		if p != nil {
			*out = append(*out, p.Val)
			seen[p] = true
			parent[p.Right] = p
			n = p.Right
		} else {
			n = nil
		}

	}

	return *out
}

func InorderTraversal(root *TreeNode) []int {
	out := &[]int{}

	visit(root, out)

	return *out
}

func visit(n *TreeNode, c *[]int) {
	if n == nil {
		return
	}

	visit(n.Left, c)
	*c = append(*c, n.Val)
	visit(n.Right, c)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || isMirrorTree(root.Left, root.Right)
}

func isMirrorTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) || (p != nil && q != nil && p.Val == q.Val && isMirrorTree(p.Left, q.Right) && isMirrorTree(p.Right, q.Left))
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createTree(nums, 0, len(nums)-1)
}

func unsortedArrayToBST(nums []int) *TreeNode {
	sort.Ints(nums)
	return sortedArrayToBST(nums)
}

func createTree(nums []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  createTree(nums, l, mid-1),
		Right: createTree(nums, mid+1, r),
	}
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func binaryTreePaths(n *TreeNode, path ...string) (out []string) {
	path = append(path, fmt.Sprintf("%v", n.Val))

	if n.Left == nil && n.Right == nil {
		return []string{strings.Join(path, "->")}
	}

	if n.Left != nil {
		out = append(out, binaryTreePaths(n.Left, path...)...)
	}

	if n.Right != nil {
		out = append(out, binaryTreePaths(n.Right, path...)...)
	}

	return out
}

func binaryTreePaths2(root *TreeNode) []string {
	slice := &[]string{}
	return *binaryTreeAppendPath(slice, "", root)
}

func binaryTreeAppendPath(slice *[]string, pfx string, n *TreeNode) *[]string {
	if n == nil {
		return slice
	}

	if pfx != "" {
		pfx += "->"
	}

	pfx += strconv.Itoa(n.Val)

	if n.Left == nil && n.Right == nil {
		*slice = append(*slice, pfx)
		return slice
	}

	binaryTreeAppendPath(slice, pfx, n.Left)
	return binaryTreeAppendPath(slice, pfx, n.Right)
}
