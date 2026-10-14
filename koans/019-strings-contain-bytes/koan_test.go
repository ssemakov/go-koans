package main

// Strings contain bytes
// Fill the byte lengths and indexed byte value. Go string indexing reads a byte, even for
// UTF-8 text.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ascii := "Go"
	accent := "é"
	mixed := "é世"

	wantASCIIBytes := -999 // TODO: replace this value.
	if got := len(ascii); got != wantASCIIBytes {
		t.Errorf("ASCIIBytes: got %#v; want %#v", got, wantASCIIBytes)
	}

	wantAccentBytes := -999 // TODO: replace this value.
	if got := len(accent); got != wantAccentBytes {
		t.Errorf("AccentBytes: got %#v; want %#v", got, wantAccentBytes)
	}

	wantMixedBytes := -999 // TODO: replace this value.
	if got := len(mixed); got != wantMixedBytes {
		t.Errorf("MixedBytes: got %#v; want %#v", got, wantMixedBytes)
	}

	wantFirstByte := -999 // TODO: replace this value.
	if got := int(accent[0]); got != wantFirstByte {
		t.Errorf("FirstByte: got %#v; want %#v", got, wantFirstByte)
	}
}
