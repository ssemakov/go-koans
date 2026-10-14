package main

// Range gives index and value
// Implement positions with range. Return the indices where the target occurs, not the
// matching values. An absent target returns no indices.

import (
	"slices"
	"testing"
)

func positions(values []int, target int) []int {
	return nil // TODO: collect the matching indices.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		target int
		want   []int
	}{
		{nil, 7, nil}, {[]int{7, 2, 7}, 7, []int{0, 2}}, {[]int{2, 7}, 7, []int{1}},
		{[]int{1, 2}, 7, nil}, {[]int{0, 0}, 0, []int{0, 1}},
	}
	for _, tc := range cases {
		if got := positions(tc.values, tc.target); !slices.Equal(got, tc.want) {
			t.Errorf("positions(%v,%d)=%v; want %v", tc.values, tc.target, got, tc.want)
		}
	}
}
