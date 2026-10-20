package main

// Pack and unpack bit fields
// Implement packFields and unpackFields using shifts and masks. A nibble is four bits:
// pack the low four bits of high into bits 7-4 and the low four bits of low into bits 3-0.
// Unpack both fields as values from 0 through 15. Keep the checks.

import (
	"testing"
)

func packFields(high, low uint8) uint8 {
	return 0 // TODO: mask and pack both fields.
}
func unpackFields(packed uint8) (uint8, uint8) {
	return 0, 0 // TODO: extract the high and low nibbles.
}

func TestKoan(t *testing.T) {
	packedCases := []struct{ high, low, want uint8 }{
		{0, 0, 0}, {1, 2, 0x12}, {0x0a, 0x05, 0xa5}, {0x0f, 0x0f, 0xff},
		{0x12, 0x34, 0x24}, {0xff, 0xff, 0xff}, {0x80, 0x08, 0x08},
	}
	for _, tc := range packedCases {
		if got := packFields(tc.high, tc.low); got != tc.want {
			t.Errorf("packFields(%08b,%08b)=%08b; want %08b", tc.high, tc.low, got, tc.want)
		}
	}
	unpackedCases := []struct{ packed, high, low uint8 }{
		{0, 0, 0}, {0x12, 1, 2}, {0xa5, 0x0a, 0x05}, {0xf0, 0x0f, 0},
		{0x0f, 0, 0x0f}, {0xff, 0x0f, 0x0f}, {0x80, 8, 0},
	}
	for _, tc := range unpackedCases {
		high, low := unpackFields(tc.packed)
		if high != tc.high || low != tc.low {
			t.Errorf("unpackFields(%08b)=(%04b,%04b); want (%04b,%04b)", tc.packed, high, low, tc.high, tc.low)
		}
	}
	for value := 0; value < 256; value++ {
		packed := uint8(value)
		high, low := unpackFields(packed)
		if high > 15 || low > 15 {
			t.Fatalf("fields must fit in four bits: high=%d low=%d", high, low)
		}
		if got := packFields(high, low); got != packed {
			t.Fatalf("round trip for %08b: got %08b", packed, got)
		}
	}
}
