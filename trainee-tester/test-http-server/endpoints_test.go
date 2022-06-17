package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingHandler(t *testing.T) {
	// 1
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	// 2
	wr := httptest.NewRecorder()
	pingHandler(wr, req)

	// 3
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "pong") {
		t.Errorf(
			`response body "%s" does not contain "NAME"`,
			wr.Body.String(),
		)
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Vadim", nil)

	wr := httptest.NewRecorder()
	helloHandler(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Vadim") {
		t.Errorf(
			`response body "%s" does not contain "NAME"`,
			wr.Body.String(),
		)
	}
}

// func TestTimeoutHandler(t *testing.T) {
// 	wr := httptest.NewRecorder()
// 	req := httptest.NewRequest(http.MethodGet, "/timeout", nil)

// 	timeoutHandler(wr, req)
// 	if wr.Code != http.StatusOK {
// 		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
// 	}

// 	if !strings.Contains(wr.Body.String(), "NAME_timeout") {
// 		t.Errorf(
// 			`response body "%s" does not contain "NAME_timeout"`,
// 			wr.Body.String(),
// 		)
// 	}
// }

// func BenchmarkTimeoutHandler(b *testing.B) {
// 	wr := httptest.NewRecorder()
// 	req := httptest.NewRequest(http.MethodGet, "/timeout", nil)

// 	timeoutHandler(wr, req)
// 	if wr.Code != http.StatusOK {
// 		b.Errorf("got HTTP status code %d, expected 200", wr.Code)
// 	}

// 	if !strings.Contains(wr.Body.String(), "NAME_timeout") {
// 		b.Errorf(
// 			`response body "%s" does not contain "NAME_timeout"`,
// 			wr.Body.String(),
// 		)
// 	}
// }
