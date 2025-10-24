package convert

import "time"

func ToUTCString(localTime, layout, timeZone string) (string, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}
	t, err := time.ParseInLocation(layout, localTime, loc)
	if err != nil {
		return "", err
	}
	return t.UTC().Format(layout), nil
}
