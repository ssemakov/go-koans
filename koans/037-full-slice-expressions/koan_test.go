package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := []int{1, 2, 3, 4}
	part := values[1:2:3]
	want := -999 // TODO: replace this value.
	got := cap(part)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
