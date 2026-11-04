package main

// A function is a value
// Complete the call to fn inside transform. A function can be passed as an argument
// and called through its parameter name. Call it once for each value in order; the
// slice allocation and loop are provided. Keep the checks.

import (
	"slices"
	"testing"
)

func double(value int) int { return value * 2 }
func negate(value int) int { return -value }

func transform(values []int, fn func(int) int) []int {
	result := make([]int, len(values))
	for index, value := range values {
		result[index] = value // TODO: call fn with value and store its result.
	}
	return result
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		fn     func(int) int
		want   []int
	}{
		{nil, double, nil},
		{[]int{1, 2, 3}, double, []int{2, 4, 6}},
		{[]int{1, 2, 3}, negate, []int{-1, -2, -3}},
	}
	for _, tc := range cases {
		before := slices.Clone(tc.values)
		var visited []int
		got := transform(tc.values, func(n int) int { visited = append(visited, n); return tc.fn(n) })
		if !slices.Equal(got, tc.want) {
			t.Errorf("transform(%v)=%v; want %v", before, got, tc.want)
		}
		if !slices.Equal(visited, before) {
			t.Errorf("callback inputs=%v; want %v in order", visited, before)
		}
		if len(got) > 0 {
			got[0] = 99
		}
		if !slices.Equal(tc.values, before) {
			t.Error("input changed")
		}
	}
}
