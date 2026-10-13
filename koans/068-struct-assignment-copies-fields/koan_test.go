package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	type Point struct{ X int }
	a := Point{X: 3}
	b := a
	b.X = 5
	want := -999 // TODO: replace this value.
	got := a.X
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
