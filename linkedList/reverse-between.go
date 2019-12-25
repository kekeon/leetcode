package linkedList

import "fmt"

/**
 92. 反转链表 II
 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 输出: 1->4->3->2->5->NULL
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: head,
	}

	var h *ListNode

	node := res

	for i := 1; i < m; i++ {
		node = node.Next
		h = node.Next
	}

	nextHead := node.Next

	var next *ListNode
	var pre *ListNode

/*  next := &ListNode{
		Val:  0,
		Next: nil,
	}*/

/*	pre := &ListNode{
		Val:  0,
		Next: nil,
	}*/

	for j := m; j <= n; j++ {
		next = nextHead.Next
		nextHead.Next = pre
		pre = nextHead
		nextHead = next
	}
	res.Next.Next = pre
	h.Next = *&nextHead
	fmt.Println("======")
	fmt.Println(pre)
	fmt.Println(&pre)
	fmt.Println(*pre)
	fmt.Println("end")
	res.Next.Next = *&pre

	return  res.Next
}
