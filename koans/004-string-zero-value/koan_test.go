package main

// String zero value
// Fill the expected strings. An empty input and an omitted field do not supply a default
// name automatically.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var name string
	record := struct{ Name string }{}
	initialized := struct{ Name string }{Name: "Ada"}

	wantVariable := "TODO" // TODO: replace this value.
	if got := name; got != wantVariable {
		t.Errorf("Variable: got %#v; want %#v", got, wantVariable)
	}

	wantOmittedField := "TODO" // TODO: replace this value.
	if got := record.Name; got != wantOmittedField {
		t.Errorf("OmittedField: got %#v; want %#v", got, wantOmittedField)
	}

	wantInitializedField := "TODO" // TODO: replace this value.
	if got := initialized.Name; got != wantInitializedField {
		t.Errorf("InitializedField: got %#v; want %#v", got, wantInitializedField)
	}
}
