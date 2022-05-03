package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/welcome", nil)

	messageHandler(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Welcome to Go Web Development") {
		t.Errorf(
			`response body "%s" does not contain "Welcome to Go Web Development"`,
			wr.Body.String(),
		)
	}
}
