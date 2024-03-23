package listnode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	v := make([]int, 0, 50000)

	h := head
	for h != nil {
		v = append(v, h.Val)
		h = h.Next
	}

	h = head
	l, r := 0, len(v)-1
	left := true
	for l <= r {
		if left {
			h.Val = v[l]
			l++
		} else {
			h.Val = v[r]
			r--
		}
		left = !left
		h = h.Next
	}
}
