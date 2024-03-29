package binarytree

import "fmt"

type BN[T any] struct {
	lp     int
	rp     int
	parent *BN[T]
	left   *BN[T]
	right  *BN[T]
	val    T
}

func RootBN[T any](t T) *BN[T] {
	return &BN[T]{val: t, lp: 1, rp: 1}
}

func (n *BN[T]) pad(c *BN[T]) {
	if n.left == c {
		n.lp += 2
	} else {
		n.rp += 2
	}

	if n.parent != nil {
		n.parent.pad(n)
	}
}

func (n *BN[T]) addToLeft(l *BN[T]) {
	if l == nil {
		return
	}
	n.left = l
	l.parent = n
	// r.parent.pad(l)
	l.parent.lp = l.lp + l.rp - 1
}

func (n *BN[T]) addToRight(r *BN[T]) {
	if r == nil {
		return
	}
	n.right = r
	r.parent = n
	// r.parent.pad(r)
	r.parent.rp = r.lp + r.rp
}

func T2N(t *TreeNode) *BN[*TreeNode] {
	b := RootBN(t)
	if t.Left != nil {
		l := T2N(t.Left)
		l.rp++
		l.lp++
		l.lp++
		// l.rp++
		// l.rp++

		b.addToLeft(l)
	}

	if t.Right != nil {
		r := T2N(t.Right)
		r.lp++
		r.rp++
		r.rp++
		b.addToRight(r)
	}

	return b
}

func PrintBN(n *BN[*TreeNode]) {
	if n == nil {
		return
	}

	l1 := []*BN[*TreeNode]{n}
	l2 := []*BN[*TreeNode]{}

	w := 1
	for w > 0 {
		for _, b := range l1 {
			w--
			for i := 0; i < b.lp; i++ {
				fmt.Print("  ")
			}

			fmt.Print(b.val.Val)

			for i := 0; i < b.rp; i++ {
				fmt.Print("  ")
			}

			if b.left != nil {
				l2 = append(l2, b.left)
				w++
			}
			if b.right != nil {
				l2 = append(l2, b.right)
				w++
			}
		}
		fmt.Println()
		fmt.Println()
		l1 = l2
		l2 = []*BN[*TreeNode]{}
	}

}
