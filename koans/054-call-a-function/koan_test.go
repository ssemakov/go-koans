package main

// Call a function
// Implement clamp. Return the nearest bound when value is outside lower..upper, and the
// value itself otherwise. lower is never greater than upper.

import (
	"testing"
)

func clamp(value, lower, upper int) int {
	return value // TODO: implement the bounds.
}

func TestKoan(t *testing.T) {
	cases := []struct{ value, lower, upper, want int }{
		{-1, 0, 10, 0}, {0, 0, 10, 0}, {4, 0, 10, 4}, {10, 0, 10, 10}, {11, 0, 10, 10},
		{-7, -5, -2, -5}, {-3, -5, -2, -3}, {9, 4, 4, 4},
	}
	for _, tc := range cases {
		if got := clamp(tc.value, tc.lower, tc.upper); got != tc.want {
			t.Errorf("clamp(%d,%d,%d)=%d; want %d", tc.value, tc.lower, tc.upper, got, tc.want)
		}
	}
}
