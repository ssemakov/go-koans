package main

// Struct fields
// Implement newPerson with a struct literal. Store the supplied name and age; leave Active
// at its zero value.

import (
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Active bool
}

func newPerson(name string, age int) Person {
	return Person{} // TODO: initialize the supplied fields.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		name string
		age  int
	}{{"Ada", 30}, {"Grace", 42}, {"", 0}}
	for _, tc := range cases {
		got := newPerson(tc.name, tc.age)
		want := Person{Name: tc.name, Age: tc.age}
		if got != want {
			t.Errorf("newPerson(%q,%d)=%+v; want %+v", tc.name, tc.age, got, want)
		}
		if got.Active {
			t.Error("Active must default to false")
		}
	}
}
