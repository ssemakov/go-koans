package main

import (
	"testing"
)

type Speaker interface{ Speak() string }
type Greeting string

func (g Greeting) Speak() string { return string(g) }
func TestKoan(t *testing.T) {
	var value Speaker = Greeting("hello")
	want := "__" // TODO: replace this value.
	got := value.Speak()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
