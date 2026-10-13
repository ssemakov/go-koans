package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := make([]int, 3, 8)
	want := -999 // TODO: replace this value.
	got := cap(value)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
