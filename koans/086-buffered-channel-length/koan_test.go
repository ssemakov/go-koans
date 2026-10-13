package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 7
	want := -999 // TODO: replace this value.
	got := len(ch)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
