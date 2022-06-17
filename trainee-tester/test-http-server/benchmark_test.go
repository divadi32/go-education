package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkPingHandler(b *testing.B) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	pingHandler(wr, req)
}

func BenchmarkHelloHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/hello?name=NAME", nil)

		helloHandler(wr, req)
	}
}
