package helper

import (
	"pendaftaran-pasien-backend/internal/custom"
	"strings"
	"time"
)

func ConvertDayToEnglish(day string) (string, error) {
	dayMap := map[string]string{
		"senin":  "Monday",
		"selasa": "Tuesday",
		"rabu":   "Wednesday",
		"kamis":  "Thursday",
		"jumat":  "Friday",
		"sabtu":  "Saturday",
		"minggu": "Sunday",
	}

	dayLower := strings.ToLower(day)

	if dayEng, exists := dayMap[dayLower]; exists {
		return strings.ToLower(dayEng), nil
	}

	return "", custom.ErrNotFound
}

func ConvertTimeToDay(time time.Time) string {
	return time.Format("Monday")
}
