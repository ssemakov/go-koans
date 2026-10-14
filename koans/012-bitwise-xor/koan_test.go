package main

// Bitwise XOR
// Implement toggle using XOR. Change only the bits selected by mask; applying the same
// mask twice must restore the value.

import (
	"testing"
)

func toggle(value, mask uint8) uint8 {
	return value // TODO: implement.
}

func TestKoan(t *testing.T) {
	cases := []struct{ value, mask, want uint8 }{
		{5, 2, 7}, {7, 2, 5}, {5, 1, 4}, {0, 8, 8}, {255, 0, 255},
	}
	for _, tc := range cases {
		got := toggle(tc.value, tc.mask)
		if got != tc.want {
			t.Errorf("toggle(%08b,%08b)=%08b; want %08b", tc.value, tc.mask, got, tc.want)
		}
		if restored := toggle(got, tc.mask); restored != tc.value {
			t.Errorf("toggling twice gave %d; want %d", restored, tc.value)
		}
	}
}
