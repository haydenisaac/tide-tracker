package tidetimes

import (
	"errors"
	"testing"
)

func TestParseText(t *testing.T) {

	tests := []struct {
		name        string
		text        string
		want        float64
		expectedErr error
	}{
		{"success", "approximately 4.35m", 4.35, nil},
		{"no number", "approx", -1, ErrUnableToParse},
		{"double digit", "approximately 10.00m", 10.00, nil},
		{"single decimal", "approximately 10.5m", 10.5, nil},
		{"two numbers", "approximately 10.0m,  max height 15.0", 10.0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseText(tt.text)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.expectedErr)
			}

			if got != tt.want {
				t.Errorf("unexpected number: got %g, want %g", got, tt.want)
			}
		})
	}

}
