package main

// Implicit interface implementation
// Complete Robot.Name so Robot satisfies Named's behavior. The same consumer must accept
// people and robots; robot names are robot: followed by the ID.

import (
	"slices"
	"testing"
)

type Named interface{ Name() string }
type Person struct{ Value string }

func (p Person) Name() string { return p.Value }

type Robot struct{ ID string }

func (r Robot) Name() string {
	return "" // TODO: implement the interface method.
}
func greetings(values []Named) []string {
	var result []string
	for _, value := range values {
		result = append(result, "Hello "+value.Name())
	}
	return result
}

func TestKoan(t *testing.T) {
	values := []Named{Person{Value: "Ada"}, Robot{ID: "r1"}, Person{Value: "Grace"}, Robot{ID: "r2"}}
	want := []string{"Hello Ada", "Hello robot:r1", "Hello Grace", "Hello robot:r2"}
	if got := greetings(values); !slices.Equal(got, want) {
		t.Errorf("greetings=%q; want %q", got, want)
	}
}
