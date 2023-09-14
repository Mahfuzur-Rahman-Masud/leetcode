package leet

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	out := &[]int{}
	preorderRecurse(root, out)

	return *out
}

func preorderRecurse(root *Node, out *[]int) {
	if root == nil {
		return
	}

	*out = append(*out, root.Val)
	for _, child := range root.Children {
		preorderRecurse(child, out)
	}
}

func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}

	depth := 0
	for _, c := range node.Children {
		d := maxDepth(c)
		if d > depth {
			depth = d
		}
	}
	return depth + 1
}

func NewTree(nodes []int) *Node {
	parents := []*Node{}
	layer := []*Node{}
	p := 0

	var root *Node

	for _, n := range nodes {
		lp := len(parents)
		if n >= 0 {
			node := &Node{Val: n, Children: []*Node{}}
			if root == nil {
				root = node
			}
			layer = append(layer, node)
			if p < lp {
				parent := parents[p]
				parent.Children = append(parent.Children, node)
				p++
			} else {
				p = 0
				parents = layer
				layer = []*Node{}
			}
		} else {
			p++
			if p >= lp {
				p = 0
				parents = layer
				layer = []*Node{}
			}
		}
	}

	return root
}

func maxDepthRecurse(node *Node) int {
	if node == nil {
		return 0
	}

	depth := 0
	for _, c := range node.Children {
		d := maxDepthRecurse(c)
		if d > depth {
			depth = d
		}
	}
	return depth + 1
}
