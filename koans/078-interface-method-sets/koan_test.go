package main

import (
	"testing"
)

type Incrementer interface{ Increment() }
type Counter struct{ N int }

func (c *Counter) Increment() { c.N++ }
func TestKoan(t *testing.T) {
	var value Incrementer = &Counter{}
	value.Increment()
	want := -999 // TODO: replace this value.
	got := value.(*Counter).N
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
