package linkedList

/**
 92. 反转链表 II
 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 输出: 1->4->3->2->5->NULL
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if  m == n {
		return head
	}

	res := &ListNode{
		Val:  0,
		Next: head,
	}

	node := res

	for i := 1; i < m; i++ {
		node = node.Next
	}

	nextHead := node.Next

	var next *ListNode
	var pre *ListNode


	for j := m; j <= n; j++ {
		next = nextHead.Next
		nextHead.Next = pre
		pre = nextHead
		nextHead = next
	}
	node.Next.Next = next
	node.Next = pre

	return  res.Next
}
