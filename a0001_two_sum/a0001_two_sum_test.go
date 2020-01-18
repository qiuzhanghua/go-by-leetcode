package a0001_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	expected := []int{0, 1}
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}

func TestTwoSum_2(t *testing.T) {
	expected := []int{1, 2}
	actual := twoSum([]int{9, 7, 2, 15}, 9)
	eq := reflect.DeepEqual(actual, expected)
	if !eq {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}
