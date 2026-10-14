package main

// Type parameters
// Implement generic first. Return the first value and true, or the type's zero value and
// false for empty input. The same function must handle numbers, strings, and slices.

import (
	"slices"
	"testing"
)

func first[T any](values []T) (T, bool) {
	var zero T
	return zero, false // TODO: handle both empty and nonempty input.
}

func TestKoan(t *testing.T) {
	for _, tc := range []struct {
		values []int
		want   int
		ok     bool
	}{{nil, 0, false}, {[]int{0, 7}, 0, true}, {[]int{7, 8}, 7, true}} {
		got, ok := first(tc.values)
		if got != tc.want || ok != tc.ok {
			t.Errorf("first(%v)=(%d,%v)", tc.values, got, ok)
		}
	}
	for _, tc := range []struct {
		values []string
		want   string
		ok     bool
	}{{nil, "", false}, {[]string{"", "Go"}, "", true}, {[]string{"Go", "code"}, "Go", true}} {
		got, ok := first(tc.values)
		if got != tc.want || ok != tc.ok {
			t.Errorf("first(%q)=(%q,%v)", tc.values, got, ok)
		}
	}
	got, ok := first([][]int{{1, 2}, {3}})
	if !ok || !slices.Equal(got, []int{1, 2}) {
		t.Errorf("first of slices=(%v,%v)", got, ok)
	}
}
