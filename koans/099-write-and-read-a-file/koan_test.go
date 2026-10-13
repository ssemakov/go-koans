package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKoan(t *testing.T) {
	name := filepath.Join(t.TempDir(), "sample.txt")
	if err := os.WriteFile(name, []byte("Go"), 0600); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	want := "__" // TODO: replace this value.
	got := string(data)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
