package linkedList

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fastNext := head.Next
	slowNext := head

	for fastNext != slowNext {
		if fastNext == nil || fastNext.Next == nil {
			return false
		}

		fastNext = fastNext.Next.Next
		slowNext = slowNext.Next
	}
	return true

}
