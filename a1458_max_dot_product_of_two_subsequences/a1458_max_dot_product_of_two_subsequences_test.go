package a1458_max_dot_product_of_two_subsequences

import "testing"

func Test_MaxDot_01(t *testing.T) {
	nums1 := []int{2, 1, -2, 5}
	nums2 := []int{3, 0, -6}

	expected := 18
	actual := maxDotProduct(nums1, nums2)
	if expected != actual {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

}
