package main

// Read binary masks
// Fill the decimal values of these uint8 binary masks. The rightmost bit represents 1, and
// each position to its left doubles the value. Keep the binary literals and checks.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	wantRead := uint8(0) // TODO: replace this value.
	if got := uint8(0b00000001); got != wantRead {
		t.Errorf("Read: got %#v; want %#v", got, wantRead)
	}

	wantWrite := uint8(0) // TODO: replace this value.
	if got := uint8(0b00000010); got != wantWrite {
		t.Errorf("Write: got %#v; want %#v", got, wantWrite)
	}

	wantExecute := uint8(0) // TODO: replace this value.
	if got := uint8(0b00000100); got != wantExecute {
		t.Errorf("Execute: got %#v; want %#v", got, wantExecute)
	}

	wantCombined := uint8(0) // TODO: replace this value.
	if got := uint8(0b00000101); got != wantCombined {
		t.Errorf("Combined: got %#v; want %#v", got, wantCombined)
	}

	wantHighBit := uint8(0) // TODO: replace this value.
	if got := uint8(0b10000000); got != wantHighBit {
		t.Errorf("HighBit: got %#v; want %#v", got, wantHighBit)
	}

	wantAllBits := uint8(0) // TODO: replace this value.
	if got := uint8(0b11111111); got != wantAllBits {
		t.Errorf("AllBits: got %#v; want %#v", got, wantAllBits)
	}
}
