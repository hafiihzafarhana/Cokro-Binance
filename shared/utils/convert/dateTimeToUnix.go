package convert

import (
	"fmt"
	"time"
)


func TimeToUnixMilliString(timeStr, timeZone string) string {
	const layout = "2006-01-02T15:04:05"

	if timeStr == "" {
		return ""
	}

	if timeZone == "" {
		timeZone = "UTC"
	}

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}

	// Parse waktu dengan location
	t, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		// Jika parsing gagal, coba parse sebagai UTC (untuk format +Z)
		t, err = time.Parse(time.RFC3339, timeStr)
		if err != nil {
			return "" // atau log error
		}
	}

	// Konversi ke UTC karena Binance pakai UTC
	utc := t.In(time.UTC)
	return fmt.Sprintf("%d", utc.UnixMilli())
}


