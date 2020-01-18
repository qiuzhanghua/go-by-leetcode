package a0002_add_two_numbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	eachSums := 0

	for l1 != nil || l2 != nil || eachSums != 0 {
		if l1 != nil {
			eachSums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			eachSums += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: eachSums % 10}
		eachSums /= 10
		current = current.Next
	}

	return result.Next
}
