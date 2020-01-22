package a1299_replace_elements_with_right_greatest

import (
	"reflect"
	"testing"
)

func TestReplace(t *testing.T) {
	expected := []int{18, 6, 6, 6, 1, -1}
	actual := replaceElements([]int{17, 18, 5, 4, 6, 1})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
