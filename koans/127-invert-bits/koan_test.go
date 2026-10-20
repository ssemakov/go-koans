package main

// Invert bits
// Implement invertBits with unary ^. Flip all eight bits of a uint8, including leading
// zero bits. Unary ^ complements one value; binary ^ toggles bits using a second value.
// Keep the checks.

import (
	"testing"
)

func invertBits(value uint8) uint8 {
	return value // TODO: complement all eight bits.
}

func TestKoan(t *testing.T) {
	cases := []struct{ value, want uint8 }{
		{0b00000000, 0b11111111}, {0b11111111, 0b00000000},
		{0b00001111, 0b11110000}, {0b10000000, 0b01111111},
		{0b01010101, 0b10101010}, {0b11000011, 0b00111100},
	}
	for _, tc := range cases {
		got := invertBits(tc.value)
		if got != tc.want {
			t.Errorf("invertBits(%08b)=%08b; want %08b", tc.value, got, tc.want)
		}
		if restored := invertBits(got); restored != tc.value {
			t.Errorf("inverting %08b twice: got %08b; want %08b", tc.value, restored, tc.value)
		}
	}
}
