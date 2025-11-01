package convert

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToUnixMilliString(t *testing.T) {
	tests := []struct {
		name     string
		timeStr  string
		timeZone string
		expected string
	}{
		{
			name:     "valid UTC time",
			timeStr:  "2025-11-01T12:34:56",
			timeZone: "UTC",
			expected: func() string {
				tm, _ := time.ParseInLocation("2006-01-02T15:04:05", "2025-11-01T12:34:56", time.UTC)
				return fmt.Sprintf("%d", tm.UnixMilli())
			}(),
		},
		{
			name:     "valid Asia/Jakarta",
			timeStr:  "2025-11-01T19:34:56",
			timeZone: "Asia/Jakarta",
			expected: func() string {
				loc, _ := time.LoadLocation("Asia/Jakarta")
				tm, _ := time.ParseInLocation("2006-01-02T15:04:05", "2025-11-01T19:34:56", loc)
				return fmt.Sprintf("%d", tm.UTC().UnixMilli())
			}(),
		},
		{
			name:     "empty time string",
			timeStr:  "",
			timeZone: "UTC",
			expected: "",
		},
		{
			name:     "invalid timezone",
			timeStr:  "2025-11-01T12:34:56",
			timeZone: "Invalid/Zone",
			expected: func() string {
				tm, _ := time.ParseInLocation("2006-01-02T15:04:05", "2025-11-01T12:34:56", time.UTC)
				return fmt.Sprintf("%d", tm.UnixMilli())
			}(),
		},
		{
			name:     "invalid time format",
			timeStr:  "2025-11-01 12:34:56", // wrong layout
			timeZone: "UTC",
			expected: "",
		},
		{
			name:     "empty timezone defaults to UTC",
			timeStr:  "2025-11-01T12:34:56",
			timeZone: "",
			expected: func() string {
				tm, _ := time.ParseInLocation("2006-01-02T15:04:05", "2025-11-01T12:34:56", time.UTC)
				return fmt.Sprintf("%d", tm.UnixMilli())
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeToUnixMilliString(tt.timeStr, tt.timeZone)
			assert.Equal(t, tt.expected, result)
		})
	}
}