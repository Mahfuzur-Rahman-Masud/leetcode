package leet

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(values []int) *ListNode {
	if values == nil || len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}

	c := head
	for i := 1; i < len(values); i++ {
		c.Next = &ListNode{Val: values[i]}
		c = c.Next
	}

	return head
}

func compareLists(list1, list2 *ListNode) bool {
	node1, node2 := list1, list2

	for node1 != nil && node2 != nil {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}

	return node1 == nil && node2 == nil
}

func IsPalindrome2(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	// find the mid of the list
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// break the list at mid point and reverse
	var prev, curr, next *ListNode = nil, slow, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	tail := prev

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}

		head = head.Next
		tail = tail.Next
	}

	return true
}

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	c := []int{}

	current := head

	for current != nil {
		c = append(c, current.Val)
		current = current.Next
	}

	l := len(c)
	mid := l/2 + 1

	for i := 0; i < mid; i++ {
		if c[i] != c[l-i-1] {
			return false
		}
	}

	return true
}

func reverseList(head *ListNode) *ListNode {

	var prev, curr, next *ListNode = nil, head, nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func removeElements(head *ListNode, val int) *ListNode {
	c := head

	for head != nil && head.Val == val {
		head = head.Next
	}

	for c != nil {
		if c.Next != nil && c.Next.Val == val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head
}

func removeElements_iter2(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		head = head.Next
	}

	c := head

	for c != nil {
		if c.Next != nil && c.Next.Val == val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	return head
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	n1 := headA
	n2 := headB

	for n1 != n2 {
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}

		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}

	return n1
}

func hasCycle(head *ListNode) bool {
	return HasCycle(head)
}

func HasCycle(head *ListNode) bool {
	fast, slow := head, head

	// if there is a cycle first will catch up to slow

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}

func DeleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
