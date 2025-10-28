package convert

func NormalizeTimeZone(tz string) string {
	var tzMap = map[string]string{
		"Asia/Jakarta": "+7",
		"Asia/Bangkok": "+7",
		"Asia/Tokyo":   "+9",
		"Asia/Seoul":   "+9",
		"Asia/Singapore": "+8",
		"Asia/Shanghai": "+8",
		"Asia/Dubai": "+4",
		"Asia/Kolkata": "+5:30",
		"Europe/London": "0",
		"Europe/Paris": "+1",
		"America/New_York": "-5",
		"America/Los_Angeles": "-8",
		"UTC": "0",
		"":    "0", // jika kosong, default UTC
	}

	if offset, ok := tzMap[tz]; ok {
		return offset
	}

	// fallback: kalau tidak ada di map, default ke UTC
	return "0"
}