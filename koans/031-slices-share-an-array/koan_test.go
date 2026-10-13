package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := []int{1, 2, 3}
	alias := values
	alias[0] = 9
	want := -999 // TODO: replace this value.
	got := values[0]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
