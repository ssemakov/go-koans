package main

// Check bit flags
// Implement hasAny and hasAll using bitwise AND (&). hasAny reports whether at least one
// requested bit is set; hasAll requires every requested bit. An empty mask matches all,
// but does not match any. Keep the checks.

import (
	"testing"
)

func hasAny(flags, mask uint8) bool {
	return false // TODO: test whether any requested bit is set.
}
func hasAll(flags, mask uint8) bool {
	return false // TODO: test whether all requested bits are set.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		flags, mask uint8
		any, all    bool
	}{
		{0, 0, false, true}, {0, 0b001, false, false}, {0b001, 0b001, true, true},
		{0b001, 0b011, true, false}, {0b011, 0b011, true, true}, {0b100, 0b011, false, false},
		{0b10000000, 0b10000000, true, true}, {0b10000001, 0b011, true, false},
		{0xff, 0xff, true, true}, {0xff, 0, false, true},
	}
	for _, tc := range cases {
		if got := hasAny(tc.flags, tc.mask); got != tc.any {
			t.Errorf("hasAny(%08b,%08b)=%v; want %v", tc.flags, tc.mask, got, tc.any)
		}
		if got := hasAll(tc.flags, tc.mask); got != tc.all {
			t.Errorf("hasAll(%08b,%08b)=%v; want %v", tc.flags, tc.mask, got, tc.all)
		}
	}
}
