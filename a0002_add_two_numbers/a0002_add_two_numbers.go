package a0002_add_two_numbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	var head, cur *ListNode
	sum, carry := 0, 0
	for cur1 != nil && cur2 != nil {
		sum = (cur1.Val + cur2.Val + carry) % 10
		carry = (cur1.Val + cur2.Val + carry) / 10

		if head == nil {
			head = &ListNode{sum, nil}
			cur = head
		} else {
			cur.Next = &ListNode{sum, nil}
			cur = cur.Next
		}

		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	// Loop over left over l1
	for cur1 != nil {
		sum = (cur1.Val + carry) % 10
		carry = (cur1.Val + carry) / 10

		if head == nil {
			head = &ListNode{sum, nil}
			cur = head
		} else {
			cur.Next = &ListNode{sum, nil}
			cur = cur.Next
		}
		cur1 = cur1.Next
	}

	// Loop over left over l2
	for cur2 != nil {
		sum = (cur2.Val + carry) % 10
		carry = (cur2.Val + carry) / 10

		if head == nil {
			head = &ListNode{sum, nil}
			cur = head
		} else {
			cur.Next = &ListNode{sum, nil}
			cur = cur.Next
		}
		cur2 = cur2.Next
	}

	// Take in any carry
	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
		cur = cur.Next
	}
	return head

}
