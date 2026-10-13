package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var value any = "Go"
	text, ok := value.(string)
	want := false // TODO: replace this value.
	got := ok && text == "Go"
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
