package main

// For loop
// Implement sum using a for loop. Handle empty input, a single value, and negative numbers
// without changing the input.

import (
	"slices"
	"testing"
)

func sum(values []int) int {
	return 0 // TODO: accumulate with a loop.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		want   int
	}{{nil, 0}, {[]int{7}, 7}, {[]int{1, 2, 3}, 6}, {[]int{-4, 2, -1}, -3}}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		if got := sum(tc.values); got != tc.want {
			t.Errorf("sum(%v)=%d; want %d", tc.values, got, tc.want)
		}
		if !slices.Equal(tc.values, before) {
			t.Error("sum changed its input")
		}
	}
}
