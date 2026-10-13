package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var value any = 7
	result := ""
	switch value.(type) {
	case int:
		result = "int"
	case string:
		result = "string"
	}
	want := "__" // TODO: replace this value.
	got := result
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
