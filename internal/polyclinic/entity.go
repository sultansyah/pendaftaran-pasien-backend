package polyclinic

import "time"

type Polyclinic struct {
	ClinicID   string    `json:"clinic_id"`
	ClinicName string    `json:"clinic_name"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
