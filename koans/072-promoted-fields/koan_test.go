package main

// Promoted fields
// Complete newEmployee. Initialize the embedded Person through its field name, then read
// Name through both promoted and explicit selectors.

import (
	"testing"
)

type Person struct{ Name string }
type Employee struct {
	Person
	ID int
}

func newEmployee(name string, id int) Employee {
	return Employee{} // TODO: initialize the embedded value and ID.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{{"Ada", 7}, {"Grace", 9}, {"", 0}}
	for _, tc := range cases {
		employee := newEmployee(tc.name, tc.id)
		if employee.Name != tc.name || employee.Person.Name != tc.name || employee.ID != tc.id {
			t.Errorf("employee=%+v; want name=%q id=%d", employee, tc.name, tc.id)
		}
	}
}
