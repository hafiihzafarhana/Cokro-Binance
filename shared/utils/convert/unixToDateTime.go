package convert

import (
	"time"
)

func loadLocationSafe(timeZone string) *time.Location {
	if timeZone == "" {
		timeZone = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		// fallback ke UTC atau local
		loc = time.UTC
	}
	return loc
}

func UnixToDateTimeString(unixTime int64, timeZone string) string {
	loc := loadLocationSafe(timeZone)
	return time.UnixMilli(unixTime).In(loc).Format("2006-01-02 15:04:05")
}