package main

// Closure retains state
// Implement counter. Each factory call returns a closure with its own state; each call
// adds step and returns the updated value.

import (
	"slices"
	"testing"
)

func counter(start, step int) func() int {
	return func() int { return start } // TODO: retain independent state.
}

func TestKoan(t *testing.T) {
	a := counter(0, 1)
	b := counter(10, 2)
	got := []int{a(), a(), b(), a(), b()}
	want := []int{1, 2, 12, 3, 14}
	if !slices.Equal(got, want) {
		t.Errorf("interleaved counters=%v; want %v", got, want)
	}
	negative := counter(3, -2)
	if got := []int{negative(), negative()}; !slices.Equal(got, []int{1, -1}) {
		t.Errorf("negative step=%v", got)
	}
}
