package main

import (
	"net/http/httptest"
	"testing"
)

func TestKoan(t *testing.T) {
	request := httptest.NewRequest("GET", "/?name=Ada", nil)
	want := "__" // TODO: replace this value.
	got := request.URL.Query().Get("name")
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
