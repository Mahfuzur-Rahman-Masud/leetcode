package listnode

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	l1 := list1
	var tail *ListNode
	a--
	b++
	i := 0

	for l1 != nil {

		if i == a {
			l1.Next, l1 = list2, l1.Next
			i++
			continue
		}

		if i == b {
			tail = l1
			break
		}

		l1 = l1.Next
		i++
	}

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = tail

	return list1
}
