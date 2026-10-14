package main

// Slices share an array
// Fill the complete expected contents after modifying the slice view. Compare this with
// array assignment in exercise 027.

import (
	"slices"
	"testing"
)

func TestKoan(t *testing.T) {
	values := []int{1, 2, 3, 4}
	view := values[1:3]
	view[0] = 9
	wantValues := []int{-999} // TODO: all values.
	wantView := []int{-999}   // TODO: all values in the view.
	if !slices.Equal(values, wantValues) {
		t.Errorf("values=%v; want %v", values, wantValues)
	}
	if !slices.Equal(view, wantView) {
		t.Errorf("view=%v; want %v", view, wantView)
	}
}
