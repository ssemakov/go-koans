package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	a := []int{1, 2}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 9
	want := -999 // TODO: replace this value.
	got := a[0]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
