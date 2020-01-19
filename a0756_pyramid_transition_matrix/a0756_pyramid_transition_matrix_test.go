package a0756_pyramid_transition_matrix

import "testing"

func TestPyramidTransition(t *testing.T) {
	//expected := 20
	actual := pyramidTransition("BCD", []string{"BCG", "CDE", "GEA", "FFF"})
	if !actual {
		t.Errorf("Test failed, expected: true, got:  '%v'", actual)
	}
}

func TestPyramidTransition_2(t *testing.T) {
	//expected := 20
	actual := pyramidTransition("AABA", []string{"AAA", "AAB", "ABA", "ABB", "BAC"})
	if actual {
		t.Errorf("Test failed, expected: false, got:  '%v'", actual)
	}
}
