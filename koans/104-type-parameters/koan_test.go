package main

import (
	"testing"
)

func first[T any](values []T) T { return values[0] }
func TestKoan(t *testing.T) {

	want := -999 // TODO: replace this value.
	got := first([]int{7, 8})
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
