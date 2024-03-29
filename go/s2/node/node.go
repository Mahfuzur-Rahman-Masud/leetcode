package node

type Node[T comparable] struct {
	Parent      *Node[T]
	Value       T
	breadth     int
	leftOff     int
	rightOff    int
	relativePos int
	x           int
	y           int
	anchor      int

	Children []*Node[T]
}

const STANDARD_OFFSET = 4
const STANDARD_BREADTH = 6

func NewNode[T comparable](t T) *Node[T] {
	return &Node[T]{
		Value:   t,
		breadth: STANDARD_BREADTH,
	}
}

func (n *Node[T]) setParent(p *Node[T]) {
	n.Parent = p
}

func (n *Node[T]) shift(off int) {
	if off < 1 {
		return
	}

	n.leftOff += off
	n.rightOff += off
	if n.Parent != nil {
		n.Parent.shift(off)
	}
}

func (n *Node[T]) addChild(child *Node[T]) {
	n.Children = append(n.Children, child)
	child.setParent(n)
	if len(n.Children) > 1 {
		n.shift(STANDARD_OFFSET)
	}
}

func (n *Node[T]) getTotalBreadth() int {
	return n.breadth + n.leftOff + n.rightOff
}

func (n *Node[T]) AddWidth(w int) {
	n.breadth += w
	if n.Parent != nil {
		n.Parent.AddWidth(w)
	}
}
