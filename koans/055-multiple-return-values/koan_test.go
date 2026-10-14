package main

// Multiple return values
// Implement first. Return the first element and true for a nonempty slice. Return zero and
// false for empty input; a stored zero is still present.

import (
	"testing"
)

func first(values []int) (int, bool) {
	return 0, false // TODO: return both results.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		values []int
		want   int
		ok     bool
	}{{nil, 0, false}, {[]int{}, 0, false}, {[]int{0, 7}, 0, true}, {[]int{4, 5}, 4, true}}
	for _, tc := range cases {
		got, ok := first(tc.values)
		if got != tc.want || ok != tc.ok {
			t.Errorf("first(%v)=(%d,%v); want (%d,%v)", tc.values, got, ok, tc.want, tc.ok)
		}
	}
}
