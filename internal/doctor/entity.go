package doctor

import "time"

type Doctor struct {
	DoctorID       string    `json:"doctor_id"`
	ClinicID       string    `json:"clinic_id"`
	ClinicName     string    `json:"clinic_name"`
	DoctorName     string    `json:"doctor_name"`
	Specialization string    `json:"specialization"`
	Days           string    `json:"days"`
	StartTime      string    `json:"start_time"`
	EndTime        string    `json:"end_time"`
	PhoneNumber    string    `json:"phone_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
