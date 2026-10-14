package main

// Copy to independent storage
// Implement cloneInts. Preserve every element and give the result independent storage.
// Mutating either slice must leave the other unchanged; nil or empty input may return nil.

import (
	"slices"
	"testing"
)

func cloneInts(values []int) []int {
	return values // TODO: allocate and copy.
}

func TestKoan(t *testing.T) {
	for _, input := range [][]int{nil, {}, {7}, {1, 2, 3}} {
		original := slices.Clone(input)
		cloned := cloneInts(input)
		if !slices.Equal(cloned, original) {
			t.Errorf("cloneInts(%v)=%v", original, cloned)
			continue
		}
		if len(input) > 0 {
			cloned[0] = 99
			if !slices.Equal(input, original) {
				t.Error("editing clone changed input")
			}
			input[len(input)-1] = -7
			if cloned[len(cloned)-1] == -7 {
				t.Error("editing input changed clone")
			}
		}
	}
}
