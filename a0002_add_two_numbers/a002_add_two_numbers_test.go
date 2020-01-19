package a0002_add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	num5 := ListNode{
		Val:  5,
		Next: nil,
	}
	num6 := ListNode{
		Val:  6,
		Next: nil,
	}

	expected := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	actual := addTwoNumbers(&num5, &num6)

	eq := reflect.DeepEqual(*actual, expected)

	if !eq {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}
