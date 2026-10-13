package main

import (
	"testing"
)

func counter() func() int { n := 0; return func() int { n++; return n } }
func TestKoan(t *testing.T) {
	next := counter()
	next()
	want := -999 // TODO: replace this value.
	got := next()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
