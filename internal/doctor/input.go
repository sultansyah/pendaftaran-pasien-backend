package doctor

import "time"

type CreateDoctorInput struct {
	ClinicID       string    `json:"clinic_id"`
	DoctorName     string    `json:"doctor_name"`
	Specialization string    `json:"specialization"`
	Days           string    `json:"days"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	PhoneNumber    *string   `json:"phone_number"`
}

type GetDoctorInput struct {
	DoctorID string `json:"doctor_id"`
}

type GetDoctorByClinicInput struct {
	ClinicID string `json:"clinic_id"`
}
