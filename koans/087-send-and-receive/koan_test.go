package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "Go"
	want := "__" // TODO: replace this value.
	got := <-ch
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
