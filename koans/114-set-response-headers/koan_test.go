package main

import (
	"net/http/httptest"
	"testing"
)

func TestKoan(t *testing.T) {
	response := httptest.NewRecorder()
	response.Header().Set("Content-Type", "application/json")
	want := "__" // TODO: replace this value.
	got := response.Header().Get("Content-Type")
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
