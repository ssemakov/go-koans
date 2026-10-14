package main

// Write an HTTP response
// Implement createItem. POST returns status 201 and created followed by a newline. Other
// methods return 405, Allow: POST, and method not allowed followed by a newline. Set
// status before writing the body.

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createItem(writer http.ResponseWriter, request *http.Request) { // TODO: implement method handling, status, and body.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		method      string
		status      int
		body, allow string
	}{
		{"POST", http.StatusCreated, "created\n", ""},
		{"GET", http.StatusMethodNotAllowed, "method not allowed\n", "POST"},
		{"PUT", http.StatusMethodNotAllowed, "method not allowed\n", "POST"},
	}
	for _, tc := range cases {
		recorder := httptest.NewRecorder()
		createItem(recorder, httptest.NewRequest(tc.method, "/", nil))
		response := recorder.Result()
		body, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			t.Fatal(err)
		}
		if response.StatusCode != tc.status || string(body) != tc.body || response.Header.Get("Allow") != tc.allow {
			t.Errorf("%s: status=%d body=%q Allow=%q; want %d %q %q", tc.method, response.StatusCode, body, response.Header.Get("Allow"), tc.status, tc.body, tc.allow)
		}
	}
}
