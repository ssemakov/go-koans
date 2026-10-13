package main

import (
	"testing"
)

type DiceSet struct {
	Values []int
	Next   func() int
}

func (d *DiceSet) Roll(n int) { // TODO: implement.
}
func TestKoan(t *testing.T) {
	next := 0
	d := DiceSet{Next: func() int { next = next%6 + 1; return next }}
	d.Roll(3)
	if len(d.Values) != 3 {
		t.Fatalf("got %d dice; want 3", len(d.Values))
	}
	first := d.Values
	for _, v := range first {
		if v < 1 || v > 6 {
			t.Fatalf("invalid face %d", v)
		}
	}
	d.Roll(2)
	if len(d.Values) != 2 || first[0] != 1 || d.Values[0] != 4 {
		t.Fatal("new roll must have new storage and new values")
	}
	d.Roll(0)
	if len(d.Values) != 0 {
		t.Fatal("zero dice requested")
	}
}
