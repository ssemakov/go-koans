package main

import (
	"io"
	"strings"
	"testing"
)

func TestKoan(t *testing.T) {
	reader := strings.NewReader("Go")
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	want := "__" // TODO: replace this value.
	got := string(data)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
