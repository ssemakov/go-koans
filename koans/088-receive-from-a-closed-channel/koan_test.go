package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int)
	close(ch)
	value, open := <-ch
	want := false // TODO: replace this value.
	got := value == 0 && !open
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
