package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	dst := make([]int, 2)
	src := []int{1, 2, 3}
	want := -999 // TODO: replace this value.
	got := copy(dst, src)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
