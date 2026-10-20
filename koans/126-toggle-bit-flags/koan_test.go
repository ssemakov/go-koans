package main

// Toggle bit flags
// Implement toggleFlags using XOR (^). Flip each selected bit and preserve unselected
// bits. Toggling the same mask twice must restore the original flags. Keep the checks.

import (
	"testing"
)

func toggleFlags(flags, mask uint8) uint8 {
	return flags // TODO: toggle the selected bits.
}

func TestKoan(t *testing.T) {
	cases := []struct{ flags, mask, want uint8 }{
		{0, 0, 0}, {0, 0b101, 0b101}, {0b101, 0b001, 0b100},
		{0b101, 0b010, 0b111}, {0b101, 0b111, 0b010}, {0xff, 0xff, 0},
		{0b10000001, 0b10000000, 0b001}, {0b101, 0, 0b101},
	}
	for _, tc := range cases {
		got := toggleFlags(tc.flags, tc.mask)
		if got != tc.want {
			t.Errorf("toggleFlags(%08b,%08b)=%08b; want %08b", tc.flags, tc.mask, got, tc.want)
		}
		if restored := toggleFlags(got, tc.mask); restored != tc.flags {
			t.Errorf("toggling mask %08b twice: got %08b; want %08b", tc.mask, restored, tc.flags)
		}
	}
}
