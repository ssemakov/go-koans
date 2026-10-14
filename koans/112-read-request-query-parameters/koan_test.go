package main

// Read request query parameters
// Implement queryName. Read the first name query value, URL-decoded. Return false for a
// missing or empty name; repeated names use the first value.

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func queryName(request *http.Request) (string, bool) {
	return "", false // TODO: read and validate the name query value.
}

func TestKoan(t *testing.T) {
	cases := []struct {
		url, name string
		ok        bool
	}{
		{"/?name=Ada", "Ada", true}, {"/?name=Ada+Lovelace", "Ada Lovelace", true},
		{"/?name=%C3%A9", "é", true}, {"/", "", false}, {"/?name=", "", false},
		{"/?name=first&name=second", "first", true},
	}
	for _, tc := range cases {
		request := httptest.NewRequest("GET", tc.url, nil)
		name, ok := queryName(request)
		if name != tc.name || ok != tc.ok {
			t.Errorf("queryName(%q)=(%q,%v); want (%q,%v)", tc.url, name, ok, tc.name, tc.ok)
		}
	}
}
