package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := [3]int{8}
	want := -999 // TODO: replace this value.
	got := values[2]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
