package leet

type Node struct {
	Val      int
	Children []*Node
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

func NewTree(nodes []int) {
	parents := []*Node{}
	layer := []*Node{}
	p := 0

	for _, n := range nodes {
		lp := len(parents)
		if n >= 0 {
			node := &Node{Val: n, Children: []*Node{}}
			layer = append(layer, node)
			if lp > 0 {
				parent := parents[p]
				parent.Children = append(parent.Children, node)
				p++
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
