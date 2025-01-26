package doctor

type CreateDoctorInput struct {
	ClinicID       string `json:"clinic_id" binding:"required"`
	DoctorName     string `json:"doctor_name" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	Days           string `json:"days" binding:"required"`
	StartTime      string `json:"start_time" binding:"required"`
	EndTime        string `json:"end_time" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
}

type GetDoctorInput struct {
	DoctorID string `uri:"doctor_id" binding:"required"`
}

type GetDoctorByClinicInput struct {
	ClinicID string `uri:"clinic_id" binding:"required"`
}

type GetDoctorByDayAndClinicInput struct {
	Day      string `uri:"day" binding:"required"`
	ClinicID string `uri:"clinic_id" binding:"required"`
}
