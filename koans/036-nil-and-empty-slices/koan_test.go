package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var a []int
	b := []int{}
	want := false // TODO: replace this value.
	got := a == nil && b != nil
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
