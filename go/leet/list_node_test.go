package leet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicate(t *testing.T) {
	h := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3}}},
		},
	}

	i := DeleteDuplicates(h)

	fmt.Print("[")
	for i != nil {

		fmt.Print(i.Val, ", ")
		i = i.Next
	}
	fmt.Println("]")
}

func TestReverseList(t *testing.T) {

	fmt.Println("Compare result", compareLists(reverseList(MakeList([]int{1, 2})), MakeList([]int{2, 1})))

	assert.True(t, compareLists(reverseList(MakeList([]int{1, 2, 3, 4, 5})), MakeList([]int{5, 4, 3, 2, 1})))
	assert.True(t, compareLists(reverseList(MakeList([]int{1})), MakeList([]int{1})))
	assert.True(t, compareLists(reverseList(MakeList([]int{})), MakeList([]int{})))
	assert.True(t, compareLists(reverseList(MakeList([]int{1, 2})), MakeList([]int{2, 1})))
	assert.True(t, reverseList(nil) == nil)
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(MakeList([]int{1, 2, 2, 1})))
	assert.True(t, !IsPalindrome(MakeList([]int{1, 2})))
	assert.True(t, IsPalindrome2(MakeList([]int{1, 2, 2, 1})))
	assert.True(t, !IsPalindrome2(MakeList([]int{1, 2})))
	assert.True(t, IsPalindrome2(MakeList([]int{1})))
}

func BenchmarkIsPalindrome(b *testing.B) {
	IsPalindrome(MakeList([]int{1, 2, 2, 1}))
	IsPalindrome(MakeList([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	IsPalindrome(MakeList([]int{1, 2, 3, 4, 3, 2, 1}))
}

func BenchmarkIsPalindrome2(b *testing.B) {
	IsPalindrome2(MakeList([]int{1, 2, 2, 1}))
	IsPalindrome2(MakeList([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	IsPalindrome2(MakeList([]int{1, 2, 3, 4, 3, 2, 1}))
}
