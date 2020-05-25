package a1290_convert_binary_number_in_a_linked_list_to_integer

import (
	"testing"
)

func TestConvert_01(t *testing.T) {
	tail := ListNode{1, nil}
	mid := ListNode{0, &tail}
	head := ListNode{1, &mid}
	expected := 5
	actual := getDecimalValue(&head)
	if expected != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}
