package linkedList

func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return  nil
	}

	fast := head.Next.Next
	slow := head

	for fast != slow {

		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = head

	for  slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return  slow

}
