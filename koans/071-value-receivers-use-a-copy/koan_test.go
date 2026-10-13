package main

import (
	"testing"
)

type Counter struct{ N int }

func (c Counter) Increment() { c.N++ }
func TestKoan(t *testing.T) {
	value := Counter{}
	value.Increment()
	want := -999 // TODO: replace this value.
	got := value.N
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
