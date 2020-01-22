package a0092_reverse_linked_list_ii

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// from https://leetcode.com/problems/reverse-linked-list-ii/discuss/274708/golang-0ms

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	result := &ListNode{}
	result.Next = head

	preHead := result
	cur := result

	for i := 0; i < m; i++ {
		preHead = cur
		cur = cur.Next
	}

	newTail := cur // point to the final node in the reverse sub-list

	// preHead -> cur -> cur.Next
	cur = cur.Next

	// we try to reverse the sub-list
	// preHead -> oldHead ... newTail -> cur -> nextNode
	//               |           |
	//
	for i := m; i < n; i++ {
		oldHead := preHead.Next
		preHead.Next = cur

		nextNode := cur.Next
		cur.Next = oldHead

		cur = nextNode
		newTail.Next = nextNode
	}

	return result.Next
}
