package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	value := 3
	p := &value
	*p = 5
	want := -999 // TODO: replace this value.
	got := value
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
