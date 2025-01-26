package register

type CreateRegisterInput struct {
	MedicalRecordNo   string `json:"medical_record_no"`
	SessionPolyclinic string `json:"session_polyclinic"`
	ClinicID          string `json:"clinic_id"`
	DoctorID          string `json:"doctor_id"`
	Department        string `json:"department"`
	Class             string `json:"class"`
	EntryMethod       string `json:"entry_method"`
	AdmissionFee      string `json:"admission_fee"`
	MedicalProcedure  string `json:"medical_procedure"`
}

type GetRegisterInput struct {
	RegisterID string `uri:"register_id"`
}

type GetRegisterByMRNoInput struct {
	MedicalRecordNo string `uri:"medical_record_no"`
}
