package helper

import (
	"fmt"
	"time"
)

func GenerateMedicalRecordNo(total int) string {
	timeNow := time.Now()
	// Format: YYYYMMDD-HHMMSS
	timeStamp := timeNow.Format("20060102-150405")
	serialNumber := fmt.Sprintf("%03d", total) // Format with 3 digits 001, 002
	return fmt.Sprintf("%s-%s", timeStamp, serialNumber)
}
