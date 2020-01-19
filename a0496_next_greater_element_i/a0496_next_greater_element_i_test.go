package a0496_next_greater_element_i

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	expected := []int{-1, 3, -1}
	actual := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

	expected = []int{3, -1}
	actual = nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
