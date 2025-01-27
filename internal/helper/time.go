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

func ParseDate(date string) (time.Time, error) {
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func ParseToHour(hour string) (time.Time, error) {
	layout := "15:04"

	parsedTime, err := time.Parse(layout, hour)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func ParseDatetimeToDate(dateTime time.Time) (time.Time, error) {
	layout := "2006-01-02"
	dateString := dateTime.Format(layout)
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
