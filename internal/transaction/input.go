package transaction

type CreateTransactionInput struct {
	RegisterID      string  `json:"register_id"`
	RegistrationFee float64 `json:"registration_fee"`
	ExaminationFee  float64 `json:"examination_fee"`
	TotalFee        float64 `json:"total_fee"`
	Discount        float64 `json:"discount"`
	TotalPayment    float64 `json:"total_payment"`
	PaymentType     string  `json:"payment_type"`
	PaymentStatus   string  `json:"payment_status"`
}

type UpdateTransactionInput struct {
	RegistrationFee float64 `json:"registration_fee"`
	ExaminationFee  float64 `json:"examination_fee"`
	TotalFee        float64 `json:"total_fee"`
	Discount        float64 `json:"discount"`
	TotalPayment    float64 `json:"total_payment"`
	PaymentType     string  `json:"payment_type"`
	PaymentStatus   string  `json:"payment_status"`
}

type GetTransactionInput struct {
	TransactionID int `uri:"transaction_id"`
}

type GetTransactionByMedicalRecordNoInput struct {
	MedicalRecordNo string `uri:"medical_record_no"`
}
