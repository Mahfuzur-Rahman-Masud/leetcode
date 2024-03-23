package listnode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	h := head

	v := make([]int, 0, 1024)
	for h != nil {
		v = append(v, h.Val)
		h = h.Next
	}

	l, r := 0, len(v)-1

	for l < r {

		if v[l] != v[r] {
			return false
		}

		l++
		r--
	}

	return true
}
