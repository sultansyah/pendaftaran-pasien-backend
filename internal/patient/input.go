package patient

type CreatePatientInput struct {
	PatientName       string  `json:"patient_name" binding:"required"`
	Gender            string  `json:"gender" binding:"required"`
	PlaceOfBirth      string  `json:"place_of_birth" binding:"required"`
	DateOfBirth       string  `json:"date_of_birth" binding:"required"`
	Address           string  `json:"address" binding:"required"`
	PhoneNumber       string  `json:"phone_number" binding:"required"`
	IdentityType      string  `json:"identity_type" binding:"required"`
	IdentityNumber    string  `json:"identity_number" binding:"required"`
	City              string  `json:"city" binding:"required"`
	PostalCode        string  `json:"postal_code" binding:"required"`
	MedicalRecordDate string  `json:"medical_record_date" binding:"required"`
	BirthWeight       float64 `json:"birth_weight" binding:"required"`
	Ethnicity         string  `json:"ethnicity" binding:"required"`
	Subdistrict       string  `json:"subdistrict" binding:"required"`
	District          string  `json:"district" binding:"required"`
	REGency           string  `json:"regency" binding:"required"`
	Province          string  `json:"province" binding:"required"`
	Citizenship       string  `json:"citizenship" binding:"required"`
	Country           string  `json:"country" binding:"required"`
	Language          string  `json:"language" binding:"required"`
	BloodType         string  `json:"blood_type" binding:"required"`
	KKNumber          string  `json:"kk_number" binding:"required"`
	MaritalStatus     string  `json:"marital_status" binding:"required"`
	Religion          string  `json:"religion" binding:"required"`
	Occupation        string  `json:"occupation" binding:"required"`
	Education         string  `json:"education" binding:"required"`
	NPWP              string  `json:"npwp" binding:"required"`
	FileLocation      string  `json:"file_location" binding:"required"`

	RelativeName           string `json:"relative_name" binding:"required"`
	RelativeRelationship   string `json:"relative_relationship" binding:"required"`
	RelativePhone          string `json:"relative_phone" binding:"required"`
	RelativeIdentityNumber string `json:"relative_identity_number" binding:"required"`
	RelativeOccupation     string `json:"relative_occupation" binding:"required"`
	RelativeAddress        string `json:"relative_address" binding:"required"`
	RelativeCity           string `json:"relative_city" binding:"required"`
	RelativePostalCode     string `json:"relative_postal_code" binding:"required"`

	MotherMedicalRecordNo *string `json:"mother_medical_record_no" binding:"omitempty"`
}

type GetPatientInput struct {
	MedicalRecordNo string `uri:"medical_record_no" binding:"required"`
}
