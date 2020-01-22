package a0092_reverse_linked_list_ii

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: nil}}}}}
	actual := reverseBetween(list, 2, 4)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
