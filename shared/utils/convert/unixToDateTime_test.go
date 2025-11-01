package convert

import (
	"testing"
	"time"

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
			unixTime: 1700000000000, // contoh timestamp
			timeZone: "UTC",
			expected: "2023-11-14 22:13:20", // sesuaikan dengan timestamp
		},
		{
			name:     "Asia/Jakarta",
			unixTime: 1700000000000,
			timeZone: "Asia/Jakarta",
			expected: "2023-11-15 05:13:20", // UTC+7
		},
		{
			name:     "Empty timezone defaults Asia/Jakarta",
			unixTime: 1700000000000,
			timeZone: "",
			expected: "2023-11-15 05:13:20", // UTC+7
		},
		{
			name:     "Invalid timezone fallback local",
			unixTime: 1700000000000,
			timeZone: "Inva/Lid",
			expected: time.UnixMilli(1700000000000).In(time.Local).Format("2006-01-02 15:04:05"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnixToDateTimeString(tt.unixTime, tt.timeZone)
			assert.Equal(t, tt.expected, result)
		})
	}
}
