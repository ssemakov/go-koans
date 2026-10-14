package main

// Comparable type parameters
// Implement generic indexOf using ==. Return the first matching index and true; return -1
// and false if absent. comparable supports strings, numbers, and comparable structs, but
// not slices.

import (
	"testing"
)

type Point struct{ X, Y int }

func indexOf[T comparable](values []T, target T) (int, bool) {
	return 0, false // TODO: compare values and return the first index.
}

func TestKoan(t *testing.T) {
	for _, tc := range []struct {
		values        []int
		target, index int
		ok            bool
	}{
		{nil, 7, -1, false}, {[]int{7, 2, 7}, 7, 0, true}, {[]int{2, 7}, 7, 1, true}, {[]int{1, 2}, 7, -1, false},
	} {
		index, ok := indexOf(tc.values, tc.target)
		if index != tc.index || ok != tc.ok {
			t.Errorf("indexOf(%v,%d)=(%d,%v); want (%d,%v)", tc.values, tc.target, index, ok, tc.index, tc.ok)
		}
	}
	if index, ok := indexOf([]string{"Go", "code"}, "code"); index != 1 || !ok {
		t.Errorf("string index=(%d,%v)", index, ok)
	}
	if index, ok := indexOf([]Point{{1, 2}, {3, 4}}, Point{3, 4}); index != 1 || !ok {
		t.Errorf("struct index=(%d,%v)", index, ok)
	}
}
