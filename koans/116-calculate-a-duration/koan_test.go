package main

import (
	"testing"
	"time"
)

func TestKoan(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(90 * time.Minute)
	want := -999 // TODO: replace this value.
	got := int(end.Sub(start).Minutes())
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
