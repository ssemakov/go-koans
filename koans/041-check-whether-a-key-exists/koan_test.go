package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	values := map[string]int{"a": 0}
	_, exists := values["a"]
	want := false // TODO: replace this value.
	got := exists
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
