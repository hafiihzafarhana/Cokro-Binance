package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTimeZone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Asia/Jakarta", "Asia/Jakarta", "+7"},
		{"Asia/Kolkata", "Asia/Kolkata", "+5:30"},
		{"UTC", "UTC", "0"},
		{"Empty string", "", "0"},
		{"Unknown timezone", "Mars/Phobos", "0"},
		{"Europe/London", "Europe/London", "0"},
		{"America/New_York", "America/New_York", "-5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeTimeZone(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}