package main

// Range over an integer
// Implement indices using range over an integer. Return 0 through n-1; zero or negative n
// should produce no values.

import (
	"slices"
	"testing"
)

func indices(n int) []int {
	return nil // TODO: range over n.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{{-2, nil}, {0, nil}, {1, []int{0}}, {4, []int{0, 1, 2, 3}}}
	for _, tc := range cases {
		if got := indices(tc.n); !slices.Equal(got, tc.want) {
			t.Errorf("indices(%d)=%v; want %v", tc.n, got, tc.want)
		}
	}
}
