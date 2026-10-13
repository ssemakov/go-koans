package main

import (
	"testing"
)

type Person struct{ Name string }
type Employee struct {
	Person
	ID int
}

func TestKoan(t *testing.T) {
	value := Employee{Person: Person{Name: "Ada"}}
	want := "__" // TODO: replace this value.
	got := value.Name
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
