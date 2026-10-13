package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestKoan(t *testing.T) {
	var out bytes.Buffer
	_, err := io.CopyN(&out, strings.NewReader("hello world"), 5)
	if err != nil {
		t.Fatal(err)
	}
	want := "__" // TODO: replace this value.
	got := out.String()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
