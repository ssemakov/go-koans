package main

// Sort a slice
// Implement sortedCopy with slices.Sort. Return a sorted independent copy; preserve the
// input and its storage. Empty input may produce nil.

import (
	"slices"
	"testing"
)

func sortedCopy(values []int) []int {
	return values // TODO: clone before sorting.
}

func TestKoan(t *testing.T) {
	cases := []struct{ values, want []int }{
		{nil, nil}, {[]int{7}, []int{7}}, {[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{3, 1, 3, 1}, []int{1, 1, 3, 3}}, {[]int{-1, 0, -2}, []int{-2, -1, 0}},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		got := sortedCopy(tc.values)
		if !slices.Equal(got, tc.want) {
			t.Errorf("sortedCopy(%v)=%v; want %v", before, got, tc.want)
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("sorting or editing the result changed the input")
		}
	}
}
