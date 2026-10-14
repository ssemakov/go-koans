package main

// Break a loop
// Implement takeUntilZero. Use break to stop before the first zero; ignore values after
// it. Return an independent result.

import (
	"slices"
	"testing"
)

func takeUntilZero(values []int) []int {
	return nil // TODO: stop at the first zero.
}

func TestKoan(t *testing.T) {
	cases := []struct{ values, want []int }{
		{nil, nil}, {[]int{0, 1}, nil}, {[]int{2, 3, 0, 4}, []int{2, 3}},
		{[]int{2, 3}, []int{2, 3}}, {[]int{-1, 0}, []int{-1}},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		got := takeUntilZero(tc.values)
		if !slices.Equal(got, tc.want) {
			t.Errorf("takeUntilZero(%v)=%v; want %v", before, got, tc.want)
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("the result must have independent storage")
		}
	}
}
