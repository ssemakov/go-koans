package main

// Delete a slice element
// Implement deleteAt for a valid index. Keep the remaining elements in order and leave the
// input unchanged. A result with zero elements may be nil.

import (
	"slices"
	"testing"
)

func deleteAt(values []int, index int) []int {
	return values // TODO: construct an independent result.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		index  int
		want   []int
	}{
		{[]int{1, 2, 3}, 0, []int{2, 3}}, {[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}}, {[]int{7}, 0, nil},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		got := deleteAt(tc.values, tc.index)
		if !slices.Equal(got, tc.want) {
			t.Errorf("deleteAt(%v,%d)=%v; want %v", before, tc.index, got, tc.want)
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("result shares mutable storage with input")
		}
	}
}
