package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	ch := make(chan int)
	result := ""
	select {
	case <-ch:
		result = "received"
	default:
		result = "default"
	}
	want := "__" // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
