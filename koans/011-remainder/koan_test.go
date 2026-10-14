package main

// Remainder
// Implement wrapIndex for a positive size. Return an index from 0 through size-1,
// including for negative positions.

import (
	"testing"
)

func wrapIndex(position, size int) int {
	return 0 // TODO: implement.
}

func TestKoan(t *testing.T) {
	cases := []struct{ position, size, want int }{
		{0, 5, 0}, {7, 5, 2}, {5, 5, 0}, {-1, 5, 4}, {-6, 5, 4}, {12, 1, 0},
	}
	for _, tc := range cases {
		if got := wrapIndex(tc.position, tc.size); got != tc.want {
			t.Errorf("wrapIndex(%d,%d)=%d; want %d", tc.position, tc.size, got, tc.want)
		}
	}
}
