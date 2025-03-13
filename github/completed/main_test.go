package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status %d, obtenido %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := "PFFP!"
	if w.Body.String() != expectedBody {
		t.Errorf("esperado cuerpo %q, obtenido %q", expectedBody, w.Body.String())
	}
}
