package linkedList

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := false
	res := &ListNode{
		Val: 0,
		Next: nil,
	}

	resNext := res

	for true {



		if l1 != nil {
			resNext.Val = resNext.Val + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			resNext.Val = resNext.Val + l2.Val
			l2 = l2.Next
		}


		if carry {
			resNext.Val ++
			carry = false
		}

		if resNext.Val >= 10 {
			carry = true
			resNext.Val = resNext.Val - 10
		}


		if l1 != nil || l2 != nil {
			resNext.Next = &ListNode{
				Val: 0,
				Next: nil,
			}

			resNext = resNext.Next
		} else {
			if carry {
				resNext.Next = &ListNode{
					Val: 1,
					Next: nil,
				}
			}

			break
		}
	}

	return res
}
