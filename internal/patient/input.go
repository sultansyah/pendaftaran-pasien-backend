package patient

type CreatePatientInput struct {
	PatientName       string  `json:"patient_name"`
	Gender            string  `json:"gender"`
	PlaceOfBirth      string  `json:"place_of_birth"`
	DateOfBirth       string  `json:"date_of_birth"`
	Address           string  `json:"address"`
	PhoneNumber       string  `json:"phone_number"`
	IdentityType      string  `json:"identity_type"`
	IdentityNumber    string  `json:"identity_number"`
	City              string  `json:"city"`
	PostalCode        string  `json:"postal_code"`
	MedicalRecordDate string  `json:"medical_record_date"`
	BirthWeight       float64 `json:"birth_weight"`
	Ethnicity         string  `json:"ethnicity"`
	Subdistrict       string  `json:"subdistrict"`
	District          string  `json:"district"`
	REGency           string  `json:"regency"`
	Province          string  `json:"province"`
	Citizenship       string  `json:"citizenship"`
	Country           string  `json:"country"`
	Language          string  `json:"language"`
	BloodType         string  `json:"blood_type"`
	KKNumber          string  `json:"kk_number"`
	MaritalStatus     string  `json:"marital_status"`
	Religion          string  `json:"religion"`
	Occupation        string  `json:"occupation"`
	Education         string  `json:"education"`
	NPWP              string  `json:"npwp"`
	FileLocation      string  `json:"file_location"`

	RelativeName           string `json:"relative_name"`
	RelativeRelationship   string `json:"relative_relationship"`
	RelativePhone          string `json:"relative_phone"`
	RelativeIdentityNumber string `json:"relative_identity_number"`
	RelativeOccupation     string `json:"relative_occupation"`
	RelativeAddress        string `json:"relative_address"`
	RelativeCity           string `json:"relative_city"`
	RelativePostalCode     string `json:"relative_postal_code"`

	MotherMedicalRecordNo string `json:"mother_medical_record_no"`
}

type GetPatientInput struct {
	NoMR string `uri:"no_mr"`
}
