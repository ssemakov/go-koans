package main

// Boolean expressions
// Complete the guarded expression. Return true only when the slice has a first element and
// that element is positive. Empty input must not panic.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		want   bool
	}{
		{nil, false}, {[]int{}, false}, {[]int{0}, false}, {[]int{-1, 2}, false}, {[]int{3}, true},
	}
	for _, tc := range cases {
		got := false // TODO: use a short-circuit expression.
		if got != tc.want {
			t.Errorf("values=%v: got %v; want %v", tc.values, got, tc.want)
		}
	}
}
