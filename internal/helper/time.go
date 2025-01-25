package helper

import "time"

func ParseDateTimeLocal(date string) (time.Time, error) {
	layout := "2006-01-02T15:04"

	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
