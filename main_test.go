package main

import "testing"
import "net/http"
import "net/http/httptest"
import "strings"

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/wes@golang.org",
		nil,
	)
	if err != nil {
		t.Fatal("Could not create request", err)
	}

	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200; got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "gopher wes") {
		t.Errorf("unexpected body in response")
	}
}
