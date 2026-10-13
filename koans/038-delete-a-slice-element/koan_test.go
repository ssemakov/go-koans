package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := []int{1, 2, 3}
	values = append(values[:1], values[2:]...)
	want := -999 // TODO: replace this value.
	got := values[1]
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
