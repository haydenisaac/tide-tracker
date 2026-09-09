package smtp

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	e := New("test-user", "test-password")

	if e.user != "test-user" {
		t.Error("incorrect user")
	}

	if e.dialer == nil {
		t.Error("dialer not set up")
	}
}

func TestCreateMessage(t *testing.T) {
	m := createMessage("test-user", "test-recipient", "test-text")

	var buf bytes.Buffer
	if _, err := m.WriteTo(&buf); err != nil {
		t.Fatalf("error reading message: %v", err)
	}

	if h := m.GetHeader("To"); len(h) != 1 || !slices.Contains(h, "test-recipient") {
		t.Error("recipient not set correctly")
	}
	if h := m.GetHeader("From"); len(h) != 1 || !slices.Contains(h, "test-user") {
		t.Error("user not set correctly")
	}
	if !strings.Contains(buf.String(), "test-text") {
		t.Error("body text is missing")
	}

}
