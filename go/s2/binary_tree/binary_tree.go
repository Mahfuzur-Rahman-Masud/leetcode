package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Make(val ...int) *TreeNode {
	if val == nil || len(val) == 0 {
		return nil
	}

	root := &TreeNode{Val: val[0]}
	if len(val) == 1 {
		return root
	}

	nodes := []*TreeNode{root.Left, root.Right}
	mark := 0

	for _, v := range val[1:] {
		if v == -1 {
			mark++
			continue
		}

		node := &TreeNode{Val: v}
		nodes[mark] = node
		nodes = append(nodes, node.Left, node.Right)

	}
	return root
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
