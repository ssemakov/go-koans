package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	count := 0
	for range "é世" {
		count++
	}
	want := -999 // TODO: replace this value.
	got := count
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
