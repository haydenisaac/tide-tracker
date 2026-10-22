package scrape_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"tide-tracker/scrape"
)

func createTestServer(t *testing.T, status int, body string) string {
	t.Helper()
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	t.Cleanup(fakeServer.Close)
	return fakeServer.URL
}

func TestFindPhrase(t *testing.T) {
	tests := []struct {
		name     string
		textWant string
		errWant  error
		body     string
		status   int
		phrase   string
	}{
		{"found", "test data", nil, "test data", 200, "data"},
		{"not found", "", scrape.ErrPhraseNotFound, "test data", 200, "not in html"},
		{"first found", "first test data", nil, "<html><body><p>first test data</p><p>second test data</p></body></html>", 200, "data"},
		{"skip non matching", "second test data", nil, "<html><body><p>skip</p><p>second test data</p></body></html>", 200, "data"},
		{"empty body", "", scrape.ErrPhraseNotFound, "", 200, "data"},
		{"empty phrase", "", scrape.ErrPhraseIsEmpty, "test data", 200, ""},
		{"error response", "", scrape.ErrBadStatus, "", 500, "data"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := createTestServer(t, tt.status, tt.body)
			got, err := scrape.FindPhrase(url, tt.phrase)
			if !errors.Is(err, tt.errWant) {
				t.Errorf("errors differ: got %v, want %v", err, tt.errWant)
			}
			if got != tt.textWant {
				t.Errorf("text differs: got %s, want %s", got, tt.textWant)
			}
		})
	}
}
