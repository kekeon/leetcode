package linkedList

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	first := head.Next

	last := head

	for slow != first {

		if first == nil || first.Next == nil {
			return nil
		}
		last = slow
		first = first.Next.Next
		slow = slow.Next


	}

	return last

}
