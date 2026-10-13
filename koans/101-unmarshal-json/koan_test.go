package main

import (
	"encoding/json"
	"testing"
)

func TestKoan(t *testing.T) {
	var value struct {
		Age int `json:"age"`
	}
	if err := json.Unmarshal([]byte(`{"age":30}`), &value); err != nil {
		t.Fatal(err)
	}
	want := -999 // TODO: replace this value.
	got := value.Age
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
