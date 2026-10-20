package main

// Clear bit flags
// Implement clearFlags using Go's bit-clear operator (&^). Clear every bit selected by
// mask, preserving all other bits. Clearing an absent flag or an empty mask must do
// nothing. Keep the checks.

import (
	"testing"
)

func clearFlags(flags, mask uint8) uint8 {
	return flags // TODO: clear only the selected bits.
}

func TestKoan(t *testing.T) {
	cases := []struct{ flags, mask, want uint8 }{
		{0, 0xff, 0}, {0b111, 0b010, 0b101}, {0b101, 0b010, 0b101},
		{0xff, 0b00001111, 0b11110000}, {0b10000000, 0b10000000, 0},
		{0b10000001, 0b001, 0b10000000}, {0b111, 0, 0b111}, {0xff, 0xff, 0},
	}
	for _, tc := range cases {
		if got := clearFlags(tc.flags, tc.mask); got != tc.want {
			t.Errorf("clearFlags(%08b,%08b)=%08b; want %08b", tc.flags, tc.mask, got, tc.want)
		}
	}
}
