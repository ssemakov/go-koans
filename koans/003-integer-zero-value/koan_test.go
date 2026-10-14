package main

// Integer zero value
// Fill the expected values. Compare an omitted integer field with one initialized
// explicitly.

import (
	"testing"
)

func TestKoan(t *testing.T) {
	var count int
	record := struct{ Count int }{}
	initialized := struct{ Count int }{Count: 7}

	wantVariable := -999 // TODO: replace this value.
	if got := count; got != wantVariable {
		t.Errorf("Variable: got %#v; want %#v", got, wantVariable)
	}

	wantOmittedField := -999 // TODO: replace this value.
	if got := record.Count; got != wantOmittedField {
		t.Errorf("OmittedField: got %#v; want %#v", got, wantOmittedField)
	}

	wantInitializedField := -999 // TODO: replace this value.
	if got := initialized.Count; got != wantInitializedField {
		t.Errorf("InitializedField: got %#v; want %#v", got, wantInitializedField)
	}
}
