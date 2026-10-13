package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	total := 0
	for i, value := range []int{2, 4, 6} {
		total += i + value
	}
	want := -999 // TODO: replace this value.
	got := total
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
