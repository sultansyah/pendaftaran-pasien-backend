package polyclinic

type CreatePolyclinicInput struct {
	ClinicName string `json:"clinic_name" binding:"required"`
	Location   string `json:"location" binding:"required"`
}

type GetPolyclinicInput struct {
	ClinicID string `uri:"clinic_id" binding:"required"`
}
