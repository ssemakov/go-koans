package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := make([]int, 1, 3)
	alias := values[:3]
	values = append(values, 7)
	want := -999 // TODO: replace this value.
	got := alias[1]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
