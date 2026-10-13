package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	result := 0
	if n := len("Go"); n > 0 {
		result = n
	}
	want := -999 // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
