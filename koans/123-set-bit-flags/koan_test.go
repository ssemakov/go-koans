package main

// Set bit flags
// Implement setFlags using bitwise OR (|). Set every bit selected by mask and preserve the
// others. Setting an already-set flag must leave it set. Keep the checks.

import (
	"testing"
)

func setFlags(flags, mask uint8) uint8 {
	return flags // TODO: set the selected bits.
}

func TestKoan(t *testing.T) {
	cases := []struct{ flags, mask, want uint8 }{
		{0, 0, 0}, {0, 0b101, 0b101}, {0b001, 0b010, 0b011},
		{0b101, 0b001, 0b101}, {0b101, 0b101, 0b101},
		{0b10000000, 0b001, 0b10000001}, {0xff, 0b011, 0xff},
	}
	for _, tc := range cases {
		if got := setFlags(tc.flags, tc.mask); got != tc.want {
			t.Errorf("setFlags(%08b,%08b)=%08b; want %08b", tc.flags, tc.mask, got, tc.want)
		}
	}
}
