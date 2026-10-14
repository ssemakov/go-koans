package main

// Table driven tests
// Write the per-case comparison inside t.Run. The table supplies cases below, inside, and
// above the range, including both bounds. Keep clamp and the expected results; make
// failures name the input and expected result.

import (
	"testing"
)

func clamp(value int) int {
	if value < 0 {
		return 0
	}
	if value > 10 {
		return 10
	}
	return value
}

func TestKoan(t *testing.T) {
	cases := []struct {
		name        string
		value, want int
	}{
		{"below", -1, 0}, {"lower bound", 0, 0}, {"inside", 4, 4}, {"upper bound", 10, 10}, {"above", 11, 10},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := clamp(tc.value)
			t.Errorf("TODO: compare got with tc.want; got=%d", got)
		})
	}
}
