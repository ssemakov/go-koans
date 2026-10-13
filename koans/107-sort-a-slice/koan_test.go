package main

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	values := []int{3, 1, 2}
	slices.Sort(values)
	want := -999 // TODO: replace this value.
	got := values[0]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
