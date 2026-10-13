package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	result := 0
	switch {
	case false:
		result = 1
	case true:
		result = 2
	}
	want := -999 // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
