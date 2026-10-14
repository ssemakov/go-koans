package main

// String slicing
// Fill each substring. The upper bound is exclusive; omitted bounds and equal bounds
// produce different results.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	text := "gopher"

	wantMiddle := "TODO" // TODO: replace this value.
	if got := text[2:5]; got != wantMiddle {
		t.Errorf("Middle: got %#v; want %#v", got, wantMiddle)
	}

	wantPrefix := "TODO" // TODO: replace this value.
	if got := text[:2]; got != wantPrefix {
		t.Errorf("Prefix: got %#v; want %#v", got, wantPrefix)
	}

	wantSuffix := "TODO" // TODO: replace this value.
	if got := text[2:]; got != wantSuffix {
		t.Errorf("Suffix: got %#v; want %#v", got, wantSuffix)
	}

	wantEmpty := "TODO" // TODO: replace this value.
	if got := text[2:2]; got != wantEmpty {
		t.Errorf("Empty: got %#v; want %#v", got, wantEmpty)
	}

	wantWhole := "TODO" // TODO: replace this value.
	if got := text[:]; got != wantWhole {
		t.Errorf("Whole: got %#v; want %#v", got, wantWhole)
	}
}
