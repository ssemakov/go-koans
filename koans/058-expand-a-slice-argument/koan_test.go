package main

// Expand a slice argument
// Complete addAll by expanding the slice into total's variadic arguments. Preserve the
// input and include base even for an empty slice.

import (
	"slices"
	"testing"
)

func total(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
func addAll(base int, values []int) int {
	return base // TODO: forward values with slice expansion.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		base   int
		values []int
		want   int
	}{{5, nil, 5}, {1, []int{2, 3}, 6}, {0, []int{-4, 2}, -2}}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		if got := addAll(tc.base, tc.values); got != tc.want {
			t.Errorf("addAll(%d,%v)=%d; want %d", tc.base, tc.values, got, tc.want)
		}
		if !slices.Equal(tc.values, before) {
			t.Error("input changed")
		}
	}
}
