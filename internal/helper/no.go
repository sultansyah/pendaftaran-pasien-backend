package helper

import (
	"fmt"
	"time"
)

func generateNo(total int) string {
	timeNow := time.Now()
	// Format: YYYYMMDD-HHMMSS
	timeStamp := timeNow.Format("20060102150405")
	serialNumber := fmt.Sprintf("%03d", total) // Format with 3 digits 001, 002

	return fmt.Sprintf("%s%s", timeStamp, serialNumber)
}

func GenerateMedicalRecordNo(total int) string {
	return generateNo(total)
}

func GenerateRegisterNo(total int) string {
	return fmt.Sprintf("RG%03d", total)
}
