package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnixToDateTimeString(t *testing.T) {
	tests := []struct {
		name     string
		unixTime int64
		timeZone string
		expected string
	}{
		{
			name:     "UTC timezone",
			unixTime: 1700000000000,
			timeZone: "UTC",
			expected: "2023-11-14 22:13:20",
		},
		{
			name:     "Asia/Jakarta",
			unixTime: 1700000000000,
			timeZone: "Asia/Jakarta",
			expected: "2023-11-15 05:13:20",
		},
		{
			name:     "Empty timezone defaults Asia/Jakarta",
			unixTime: 1700000000000,
			timeZone: "",
			expected: "2023-11-15 05:13:20",
		},
		{
			name:     "Invalid timezone fallback UTC",
			unixTime: 1700000000000,
			timeZone: "Inva/Lid",
			expected: "2023-11-14 22:13:20", // fallback to UTC
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnixToDateTimeString(tt.unixTime, tt.timeZone)
			assert.Equal(t, tt.expected, result)
		})
	}
}

