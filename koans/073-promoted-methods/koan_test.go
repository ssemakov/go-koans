package main

import (
	"testing"
)

type Person struct{}

func (Person) Name() string { return "Ada" }

type Employee struct{ Person }

func TestKoan(t *testing.T) {
	value := Employee{}
	want := "__" // TODO: replace this value.
	got := value.Name()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
