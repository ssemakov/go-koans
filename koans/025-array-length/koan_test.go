package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := [4]int{1, 2}
	want := -999 // TODO: replace this value.
	got := len(values)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
