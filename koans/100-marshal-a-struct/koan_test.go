package main

import (
	"encoding/json"
	"testing"
)

func TestKoan(t *testing.T) {
	value := struct {
		Name string `json:"name"`
	}{Name: "Ada"}
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	want := "__" // TODO: replace this value.
	got := string(data)
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
