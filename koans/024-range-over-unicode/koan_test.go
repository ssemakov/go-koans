package main

// Range over Unicode
// Fill the expected indices and runes for each string. range yields byte offsets and
// decoded runes, not one index per byte.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	cases := []struct {
		text    string
		indices []int
		runes   []rune
	}{
		{"Go", []int{-999}, []rune{'?'}},      // TODO: indices and runes.
		{"é世", []int{-999}, []rune{'?'}},      // TODO: indices and runes.
		{"e\u0301", []int{-999}, []rune{'?'}}, // TODO: indices and runes.
	}
	for _, tc := range cases {
		var indices []int
		var runes []rune
		for index, value := range tc.text {
			indices = append(indices, index)
			runes = append(runes, value)
		}
		if !slices.Equal(indices, tc.indices) {
			t.Errorf("%q: indices=%v; want %v", tc.text, indices, tc.indices)
		}
		if !slices.Equal(runes, tc.runes) {
			t.Errorf("%q: runes=%U; want %U", tc.text, runes, tc.runes)
		}
	}
}
