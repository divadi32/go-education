package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)

	helloHandler(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Hello") {
		t.Errorf(
			`response body "%s" does not contain "Hello"`,
			wr.Body.String(),
		)
	}
}

func BenchmarkHelloHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)

		helloHandler(wr, req)
		if wr.Code != http.StatusOK {
			b.Errorf("got HTTP status code %d, expected 200", wr.Code)
		}

		if !strings.Contains(wr.Body.String(), "Hello") {
			b.Errorf(
				`response body "%s" does not contain "Hello"`,
				wr.Body.String(),
			)
		}
	}
}
