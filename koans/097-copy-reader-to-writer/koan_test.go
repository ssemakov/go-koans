package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestKoan(t *testing.T) {
	var out bytes.Buffer
	_, err := io.Copy(&out, strings.NewReader("hello"))
	if err != nil {
		t.Fatal(err)
	}
	want := "__" // TODO: replace this value.
	got := out.String()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
