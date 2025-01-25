package polyclinic

type CreatePolyclinicInput struct {
	ClinicName string `json:"clinic_name"`
	Location   string `json:"location"`
}

type GetPolyclinicInput struct {
	ClinicID string `uri:"clinic_id"`
}
