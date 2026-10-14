package main

// Append elements
// Repair addValue. Return the slice from append so the caller receives its updated length,
// including when append allocates new storage.

import (
	"slices"
	"testing"
)

func addValue(values []int, value int) []int {
	values = append(values, value)
	return values[:len(values)-1] // TODO: return the updated slice.
}

func TestKoan(t *testing.T) {
	for _, initial := range [][]int{nil, {2, 3}, make([]int, 0, 4)} {
		before := slices.Clone(initial)
		got := addValue(initial, 7)
		want := append(slices.Clone(before), 7)
		if !slices.Equal(got, want) {
			t.Errorf("addValue(%v,7)=%v; want %v", before, got, want)
		}
		if len(initial) != len(before) {
			t.Error("the caller's old slice header changed")
		}
	}
}
