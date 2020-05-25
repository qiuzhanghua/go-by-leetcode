package a1290_convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	p := head
	for p != nil {
		ans = (ans << 1) + p.Val
		p = p.Next
	}
	return ans
}
