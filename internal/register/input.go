package register

type CreateRegisterInput struct {
	MedicalRecordNo   string `json:"medical_record_no" binding:"required"`
	SessionPolyclinic string `json:"session_polyclinic" binding:"required"`
	ClinicID          string `json:"clinic_id" binding:"required"`
	DoctorID          string `json:"doctor_id" binding:"required"`
	Department        string `json:"department" binding:"required"`
	Class             string `json:"class" binding:"required"`
	EntryMethod       string `json:"entry_method" binding:"required"`
	AdmissionFee      string `json:"admission_fee" binding:"required"`
	MedicalProcedure  string `json:"medical_procedure" binding:"required"`

	RegistrationFee float64 `json:"registration_fee" binding:"required"`
	ExaminationFee  float64 `json:"examination_fee" binding:"required"`
	TotalFee        float64 `json:"total_fee" binding:"required"`
	Discount        float64 `json:"discount" binding:"required"`
	TotalPayment    float64 `json:"total_payment" binding:"required"`
	PaymentType     string  `json:"payment_type" binding:"required"`
	PaymentStatus   string  `json:"payment_status" binding:"required"`
}

type GetRegisterInput struct {
	RegisterID string `uri:"register_id" binding:"required"`
}

type GetRegisterByMRNoInput struct {
	MedicalRecordNo string `uri:"medical_record_no" binding:"required"`
}
