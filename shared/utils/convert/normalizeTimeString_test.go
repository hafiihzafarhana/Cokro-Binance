package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTimeString(t *testing.T) {
	tests := []struct {
		name     string
		timeStr  string
		timeZone string
		expected string
	}{
		{
			name:     "empty string",
			timeStr:  "",
			timeZone: "UTC",
			expected: "",
		},
		{
			name:     "UTC without Z",
			timeStr:  "2025-11-01T12:34:56",
			timeZone: "UTC",
			expected: "2025-11-01T12:34:56Z",
		},
		{
			name:     "UTC already with Z",
			timeStr:  "2025-11-01T12:34:56Z",
			timeZone: "UTC",
			expected: "2025-11-01T12:34:56Z",
		},
		{
			name:     "UTC with offset +",
			timeStr:  "2025-11-01T12:34:56+07:00",
			timeZone: "UTC",
			expected: "2025-11-01T12:34:56+07:00",
		},
		{
			name:     "UTC with offset -",
			timeStr:  "2025-11-01T12:34:56-07:00",
			timeZone: "UTC",
			expected: "2025-11-01T12:34:56-07:00",
		},
		{
			name:     "non-UTC zone",
			timeStr:  "2025-11-01T19:34:56",
			timeZone: "Asia/Jakarta",
			expected: "2025-11-01T19:34:56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeTimeString(tt.timeStr, tt.timeZone)
			assert.Equal(t, tt.expected, result)
		})
	}
}