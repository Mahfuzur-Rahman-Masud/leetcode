package s1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDiameterOfBinaryTree(t *testing.B) {
	assert.Equal(t, 1, diameterOfBinaryTree(createTreeBreadthFirstTLRArray([]int{1, 2})))
	assert.Equal(t, 1, diameterOfBinaryTree(createTreeBreadthFirstTLRArray([]int{1, 2})))
	assert.Equal(t, 3, diameterOfBinaryTree(createTreeBreadthFirstTLRArray([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, 4, diameterOfBinaryTree(createTreeBreadthFirstTLRArray([]int{1, 2, 3, 4, 5, -1, -1, 6, 7, 8, 9})))
	assert.Equal(t, 5, diameterOfBinaryTree(createTreeBreadthFirstTLRArray([]int{1, 2, 3, 4, 5, -1, -1, 6, 7, 8, 9, -1, -1, -1, -1, 10})))
}

func Benchmark3FindMode(t *testing.B) {
	assert.Equal(t, []int{0, 1}, findMode2(createTreeBreadthFirstTLRArray([]int{1, 0, 1, 0, 0, 1, 1, 0})))

}

func Benchmark2FindMode(t *testing.B) {
	assert.Equal(t, []int{2}, findMode2(createTreeBreadthFirstTLRArray([]int{1, -1, 2, 2})))
	assert.Equal(t, []int{0}, findMode2(createTreeBreadthFirstTLRArray([]int{0})))
	assert.Equal(t, []int{2, 4}, findMode2(createTreeBreadthFirstTLRArray([]int{1, 2, 2, 2, 3, 3, 4, 4, 4})))
	assert.Equal(t, []int{5}, findMode2(createTreeBreadthFirstTLRArray([]int{1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5})))
}

func BenchmarkFindMode(t *testing.B) {
	assert.Equal(t, []int{2}, findMode(createTreeBreadthFirstTLRArray([]int{1, -1, 2, 2})))
	assert.Equal(t, []int{0}, findMode(createTreeBreadthFirstTLRArray([]int{0})))
	assert.Equal(t, []int{2, 4}, findMode(createTreeBreadthFirstTLRArray([]int{1, 2, 2, 2, 3, 3, 4, 4, 4})))
	assert.Equal(t, []int{5}, findMode(createTreeBreadthFirstTLRArray([]int{1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5})))
}

func TestSumOfLeftLeaves(t *testing.T) {
	fmt.Println(*createTreeBreadthFirstTLRArray([]int{3, 9, 20, -1, -1, 15, 7}))

	assert.Equal(t, 24, sumOfLeftLeaves(createTreeBreadthFirstTLRArray([]int{3, 9, 20, -1, -1, 15, 7})))
	assert.Equal(t, 4, sumOfLeftLeaves(createTreeBreadthFirstTLRArray([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, 0, sumOfLeftLeaves(createTreeBreadthFirstTLRArray([]int{1})))

	assert.Equal(t, 24, sumOfLeftLeaves2(createTreeBreadthFirstTLRArray([]int{3, 9, 20, -1, -1, 15, 7})))
	assert.Equal(t, 4, sumOfLeftLeaves2(createTreeBreadthFirstTLRArray([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, 0, sumOfLeftLeaves2(createTreeBreadthFirstTLRArray([]int{1})))
	// assert.Equal(t, 24, sumOfLeftLeaves(&TreeNode{
	// 	Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	// }))
}

func printTree(t *TreeNode) {
	s := []*TreeNode{}
	s = append(s, t)

	fmt.Print("[")

	fmt.Println("]")

}

func TestSameTree(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	r := IsSameTree(p, q)
	fmt.Printf("Expected: %t, Actual %t\n", true, r)

	p = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	q = &TreeNode{
		Val: 1,

		Right: &TreeNode{
			Val: 2,
		},
	}

	r = IsSameTree(p, q)
	fmt.Printf("Expected: %t, Actual %t\n", false, r)

}

func TestMirrorTree(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	r := IsSymmetric(root)
	fmt.Printf("Expected: %t, Actual %t\n", false, r)

}

func TestCreateFromSortedList(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4, 5}

	r := SortedArrayToBST(ints)
	fmt.Println(r)
}

func TestCountNodes(t *testing.T) {
	assert.Equal(t, CountNodes(UnsortedArrayToBST([]int{1, 2, 3, 4, 5, 6})), 6)
	assert.Equal(t, CountNodes(nil), 0)
	assert.Equal(t, CountNodes(UnsortedArrayToBST([]int{1, 2})), 2)
	assert.Equal(t, CountNodes(UnsortedArrayToBST([]int{1})), 1)
}

func TestBinaryTreePats(t *testing.T) {

	tn := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println(binaryTreePaths(tn))

	tn = &TreeNode{Val: 1}
	assert.Equal(t, binaryTreePaths(tn), []string{"1"})

}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	r := InorderTraversal(root)
	fmt.Println(r, "\t\t", InorderTraversal2(root))

	root = nil

	r = InorderTraversal(root)
	fmt.Println(r, "\t\t", InorderTraversal2(root))

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 21,
			},
			Right: &TreeNode{
				Val: 22,
			},
		},

		Right: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 23,
			},
			Right: &TreeNode{
				Val: 24,
			},
		},
	}

	r = InorderTraversal(root)
	fmt.Println(r, "\t\t", InorderTraversal2(root))
}

func BenchmarkInorderRecursive(b *testing.B) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 21,
			},
			Right: &TreeNode{
				Val: 22,
			},
		},

		Right: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 23,
			},
			Right: &TreeNode{
				Val: 24,
			},
		},
	}

	InorderTraversal(root)

	// fmt.Println(r)
}

func BenchmarkInorderImperative(b *testing.B) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 21,
			},
			Right: &TreeNode{
				Val: 22,
			},
		},

		Right: &TreeNode{
			Val: 12,
			Left: &TreeNode{
				Val: 23,
			},
			Right: &TreeNode{
				Val: 24,
			},
		},
	}

	InorderTraversal2(root)

	// fmt.Println(r)
}
