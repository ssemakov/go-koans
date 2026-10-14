package main

// Implement a dice set
// Implement Roll using Next exactly once per die. n is nonnegative and Next returns 1..6.
// Store n values in fresh storage, preserve every previous roll, and make no generator
// calls for zero dice.

import (
	"slices"
	"testing"
)

type DiceSet struct {
	Values []int
	Next   func() int
}

func (d *DiceSet) Roll(n int) { // TODO: allocate a fresh roll and call Next once per position.
}

func TestKoan(t *testing.T) {
	calls := 0
	d := DiceSet{Next: func() int { calls++; return (calls-1)%6 + 1 }}
	d.Roll(3)
	first := d.Values
	if !slices.Equal(first, []int{1, 2, 3}) || calls != 3 {
		t.Fatalf("first=%v calls=%d; want [1 2 3],3", first, calls)
	}
	d.Roll(2)
	second := d.Values
	if !slices.Equal(second, []int{4, 5}) || calls != 5 {
		t.Fatalf("second=%v calls=%d; want [4 5],5", second, calls)
	}
	d.Roll(0)
	if len(d.Values) != 0 || calls != 5 {
		t.Errorf("zero dice=%v calls=%d", d.Values, calls)
	}
	d.Roll(5)
	if !slices.Equal(d.Values, []int{6, 1, 2, 3, 4}) || calls != 10 {
		t.Errorf("third=%v calls=%d; want [6 1 2 3 4],10", d.Values, calls)
	}
	if !slices.Equal(first, []int{1, 2, 3}) || !slices.Equal(second, []int{4, 5}) {
		t.Errorf("previous rolls changed: first=%v second=%v", first, second)
	}
	fixedCalls := 0
	fixed := DiceSet{Next: func() int { fixedCalls++; return 6 }}
	fixed.Roll(4)
	if !slices.Equal(fixed.Values, []int{6, 6, 6, 6}) || fixedCalls != 4 {
		t.Errorf("must use the supplied generator: values=%v calls=%d", fixed.Values, fixedCalls)
	}
}
