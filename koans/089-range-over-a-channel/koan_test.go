package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 3
	close(ch)
	total := 0
	for n := range ch {
		total += n
	}
	want := -999 // TODO: replace this value.
	got := total
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
