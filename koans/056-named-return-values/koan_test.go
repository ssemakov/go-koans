package main

import (
	"testing"
)

func double(n int) (result int) { result = n * 2; return }
func TestKoan(t *testing.T) {

	want := -999 // TODO: replace this value.
	got := double(4)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
