package main

import (
	"testing"
)

type Count int

func (c Count) Double() int { return int(c) * 2 }
func TestKoan(t *testing.T) {
	value := Count(4)
	want := -999 // TODO: replace this value.
	got := value.Double()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
