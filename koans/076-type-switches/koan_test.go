package main

// Type switches
// Implement describeType using a type switch. Recognize nil, int, string, and []int; other
// dynamic types, including Label, return other.

import (
	"testing"
)

type Label string

func describeType(value any) string {
	return "other" // TODO: switch on the dynamic type.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		value any
		want  string
	}{{nil, "nil"}, {7, "int"}, {"Go", "string"}, {[]int{1}, "slice"}, {Label("Go"), "other"}, {3.5, "other"}}
	for _, tc := range cases {
		if got := describeType(tc.value); got != tc.want {
			t.Errorf("describeType(%T)=%q; want %q", tc.value, got, tc.want)
		}
	}
}
