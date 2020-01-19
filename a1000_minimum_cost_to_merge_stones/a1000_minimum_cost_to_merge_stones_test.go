package a1000_minimum_cost_to_merge_stones

import "testing"

func TestMergeStones(t *testing.T) {
	expected := 20
	actual := mergeStones([]int{3, 2, 4, 1}, 2)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
