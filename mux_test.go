package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	want := []byte([]byte(`{"status": "ok"}`))

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	m := NewMux()
	m.ServeHTTP(w, r)

	// w != http.Response あくまでrecorderである
	resp := w.Result()
	// deferじゃないんだ
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want status code 200, but %v", resp.StatusCode)
	}

	got, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error("failed to read response body")
	}

	if !bytes.Equal(got, want) {
		t.Errorf("want: %s, but got: %s ", want, got)
	}
}
