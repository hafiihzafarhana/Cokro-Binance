package convert

import "strings"

func NormalizeTimeString(timeStr, timeZone string) string {
	if timeStr == "" {
		return ""
	}

	switch timeZone {
	case "UTC":
		if !strings.HasSuffix(timeStr, "Z") && !strings.Contains(timeStr, "+") && !strings.Contains(timeStr, "-") {
			return timeStr + "Z"
		}
	default:
		// Untuk zona lain (Asia/Jakarta), biarkan seperti "2025-10-27T02:00:00"
	}

	return timeStr
}
