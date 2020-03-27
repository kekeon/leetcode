package linkedList

import (
	"testing"
	"fmt"
)

func TestReverseBetween(t *testing.T) {
	node5 := ListNode{
		Val:  5,
		Next: nil,
	}

	node4 := ListNode{
		Val:  4,
		Next: &node5,
	}

	node3 := ListNode{
		Val:  3,
		Next: &node4,
	}

	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}

	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}

	v := reverseBetween(&node1, 2, 4)
	for v != nil {
		fmt.Println(v)
		v = v.Next
	}
}



func TestHasCycle(t *testing.T) {
	/*	node5 := ListNode{
			Val:  5,
			Next: nil,
		}
	*/
	node4 := ListNode{
		Val:  -4,
		Next: nil,
	}

	node3 := ListNode{
		Val:  0,
		Next: &node4,
	}

	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}

	node1 := ListNode{
		Val:  3,
		Next: &node2,
	}


	node4.Next = &node2

	v := hasCycle(&node1)
	fmt.Println(v)
}
