package main

// A function is a value
// Implement transform. Apply the supplied function once to each value, in order, and
// return independent storage. The function is non-nil; empty input produces no calls.

import (
	"slices"
	"testing"
)

func transform(values []int, fn func(int) int) []int {
	return nil // TODO: call the supplied function for each value.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		fn     func(int) int
		want   []int
	}{
		{nil, func(n int) int { return n * 2 }, nil},
		{[]int{1, 2, 3}, func(n int) int { return n * 2 }, []int{2, 4, 6}},
		{[]int{1, 2, 3}, func(n int) int { return -n }, []int{-1, -2, -3}},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		calls := 0
		got := transform(tc.values, func(n int) int { calls++; return tc.fn(n) })
		if !slices.Equal(got, tc.want) {
			t.Errorf("transform(%v)=%v; want %v", before, got, tc.want)
		}
		if calls != len(tc.values) {
			t.Errorf("calls=%d; want %d", calls, len(tc.values))
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("input changed")
		}
	}
}
