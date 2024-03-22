package listnode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev, current := head, head.Next
	head.Next = nil

	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}

	return prev
}

func reverseListValueSwitch(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	c := head
	vs := make([]int, 0, 1024)

	for c != nil {
		vs = append(vs, c.Val)
		c = c.Next
	}

	c = head
	for i := len(vs) - 1; i > -1; i-- {
		c.Val = vs[i]
		c = c.Next
	}

	return head
}
