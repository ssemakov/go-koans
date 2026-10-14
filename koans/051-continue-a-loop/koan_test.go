package main

// Continue a loop
// Implement positiveValues with continue. Skip nonpositive elements, keep all later
// positive values in order, and leave the input unchanged.

import (
	"slices"
	"testing"
)

func positiveValues(values []int) []int {
	return nil // TODO: skip, rather than stop, on nonpositive values.
}

func TestKoan(t *testing.T) {
	cases := []struct{ values, want []int }{
		{nil, nil}, {[]int{0, -1}, nil}, {[]int{2, 0, 3, -4, 5}, []int{2, 3, 5}}, {[]int{7}, []int{7}},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		got := positiveValues(tc.values)
		if !slices.Equal(got, tc.want) {
			t.Errorf("positiveValues(%v)=%v; want %v", before, got, tc.want)
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("input changed")
		}
	}
}
