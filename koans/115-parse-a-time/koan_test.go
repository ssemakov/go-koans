package main

import (
	"testing"
	"time"
)

func TestKoan(t *testing.T) {
	value, err := time.Parse("2006-01-02", "2024-03-14")
	if err != nil {
		t.Fatal(err)
	}
	want := -999 // TODO: replace this value.
	got := value.Year()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
