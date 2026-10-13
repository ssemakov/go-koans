package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestKoan(t *testing.T) {
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("created"))
	})
	handler.ServeHTTP(response, httptest.NewRequest("POST", "/", nil))
	want := -999 // TODO: replace this value.
	got := response.Code
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
