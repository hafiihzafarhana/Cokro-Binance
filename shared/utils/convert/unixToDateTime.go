package convert

import (
	"time"
)

func UnixToDateTimeString(unixTime int64, timeZone string) string {
	var loc *time.Location
	var err error

	if timeZone == "" {
		// gunakan zona waktu lokal sistem (misal Asia/Jakarta)
		loc, err = time.LoadLocation("Asia/Jakarta")
		if err != nil {
			loc = time.Local
		}
	} else {
		loc, err = time.LoadLocation(timeZone)
		if err != nil {
			loc = time.Local
		}
	}

	return time.UnixMilli(unixTime).In(loc).Format("2006-01-02 15:04:05")
}