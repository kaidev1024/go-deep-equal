package deepequal

import "testing"

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	if !isSliceEqual(a, b) {
		t.Error("expect equal, but not equal")
	}
	a = []int{1, 2, 3, 4}
	b = []int{1, 2, 3}
	if isSliceEqual(a, b) {
		t.Error("expect not equal, but equal")
	}
}
