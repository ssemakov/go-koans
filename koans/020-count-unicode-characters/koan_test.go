package main

// Count Unicode code points
// Fill the rune counts. A rune is a Unicode code point; a base letter plus a combining
// mark can display as one character but contain two runes.

import (
	"testing"
	"unicode/utf8"
)

func TestKoan(t *testing.T) {
	accent := "é"
	decomposed := "e\u0301"
	mixed := "é世"

	wantAccentRunes := -999 // TODO: replace this value.
	if got := utf8.RuneCountInString(accent); got != wantAccentRunes {
		t.Errorf("AccentRunes: got %#v; want %#v", got, wantAccentRunes)
	}

	wantDecomposedRunes := -999 // TODO: replace this value.
	if got := utf8.RuneCountInString(decomposed); got != wantDecomposedRunes {
		t.Errorf("DecomposedRunes: got %#v; want %#v", got, wantDecomposedRunes)
	}

	wantDecomposedBytes := -999 // TODO: replace this value.
	if got := len(decomposed); got != wantDecomposedBytes {
		t.Errorf("DecomposedBytes: got %#v; want %#v", got, wantDecomposedBytes)
	}

	wantMixedRunes := -999 // TODO: replace this value.
	if got := utf8.RuneCountInString(mixed); got != wantMixedRunes {
		t.Errorf("MixedRunes: got %#v; want %#v", got, wantMixedRunes)
	}
}
