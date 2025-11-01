package convert

import "strings"

func NormalizeTimeString(timeStr, timeZone string) string {
	if timeStr == "" {
		return ""
	}

	if timeZone == "UTC" {
		// Sudah ada 'Z' di akhir → return apa adanya
		if strings.HasSuffix(timeStr, "Z") {
			return timeStr
		}

		// Sudah ada offset ±HH:MM di akhir → return apa adanya
		if len(timeStr) >= 6 {
			last6 := timeStr[len(timeStr)-6:]
			if (last6[0] == '+' || last6[0] == '-') && last6[3] == ':' {
				return timeStr
			}
		}

		// Tambahkan Z kalau belum ada
		return timeStr + "Z"
	}

	// Non-UTC → return apa adanya
	return timeStr
}
