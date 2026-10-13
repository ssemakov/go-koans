package main

import (
	"testing"
)

func TestKoan(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Ada", Age: 30}
	want := -999 // TODO: replace this value.
	got := p.Age
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
