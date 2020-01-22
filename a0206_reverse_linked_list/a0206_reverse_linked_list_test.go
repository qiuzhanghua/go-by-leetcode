package a0206_reverse_linked_list

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {

	actual := reverseList(nil)
	if actual != nil {
		t.Errorf("Test failed, expected: 'nil', got:  '%v'", actual)
	}

}

func TestReverseLinkedList_2(t *testing.T) {
	node3 := ListNode{
		Val:  3,
		Next: nil,
	}
	node2 := ListNode{
		Val:  2,
		Next: &node3,
	}
	node1 := ListNode{
		Val:  2,
		Next: &node2,
	}

	node6 := ListNode{
		Val:  1,
		Next: nil,
	}
	node5 := ListNode{
		Val:  2,
		Next: &node6,
	}
	node4 := ListNode{
		Val:  3,
		Next: &node5,
	}

	expected := &node4
	actual := reverseList(&node1)
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}
