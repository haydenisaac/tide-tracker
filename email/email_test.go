package email

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	m := createMessage("test-user", "test-recipient", "test-text", "test-subject")

	var buf bytes.Buffer
	if _, err := m.WriteTo(&buf); err != nil {
		t.Fatalf("writing message: %v", err)
	}

	tests := []struct {
		header string
		want   []string
	}{
		{"From", []string{"test-user"}},
		{"To", []string{"test-recipient"}},
		{"Subject", []string{"test-subject"}},
	}
	for _, tt := range tests {
		t.Run(tt.header, func(t *testing.T) {
			if got := m.GetHeader(tt.header); !slices.Equal(got, tt.want) {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

	out := buf.String()
	if !strings.Contains(out, "test-text") {
		t.Error("body text is missing")
	}
	if !strings.Contains(out, "text/html") {
		t.Error("content type is not text/html")
	}
}
