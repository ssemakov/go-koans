package main

// Set response headers
// Repair jsonHandler by setting both headers before the body commits the response. Inspect
// the actual response snapshot, not just the recorder's mutable header map.

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func jsonHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: move the header changes before the first write.
	_, _ = writer.Write([]byte("{\"ok\":true}\n"))
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Cache-Control", "no-store")
}

func TestKoan(t *testing.T) {
	recorder := httptest.NewRecorder()
	jsonHandler(recorder, httptest.NewRequest("GET", "/", nil))
	response := recorder.Result()
	defer response.Body.Close()
	if got := response.Header.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type=%q; want application/json", got)
	}
	if got := response.Header.Get("Cache-Control"); got != "no-store" {
		t.Errorf("Cache-Control=%q; want no-store", got)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK || string(body) != "{\"ok\":true}\n" {
		t.Errorf("status=%d body=%q", response.StatusCode, body)
	}
}
