package main

import "testing"
import "net/http"
import "net/http/httptest"
import "strings"

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"wes@golang.org", "gopher wes"},
		{"something", "dear something"},
	}

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
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
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response")
		}
	}
}
