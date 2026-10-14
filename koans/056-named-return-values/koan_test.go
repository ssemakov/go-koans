package main

// Named return values
// Complete the named result assignments. Return a quotient and remainder with a bare
// return; divisor is nonzero.

import (
	"testing"
)

func divide(value, divisor int) (quotient, remainder int) {
	// TODO: assign quotient and remainder before the bare return.
	return
}

func TestKoan(t *testing.T) {
	cases := []struct{ value, divisor, quotient, remainder int }{{7, 2, 3, 1}, {-7, 2, -3, -1}, {8, 4, 2, 0}, {0, 3, 0, 0}}
	for _, tc := range cases {
		quotient, remainder := divide(tc.value, tc.divisor)
		if quotient != tc.quotient || remainder != tc.remainder {
			t.Errorf("divide(%d,%d)=(%d,%d); want (%d,%d)", tc.value, tc.divisor, quotient, remainder, tc.quotient, tc.remainder)
		}
	}
}
