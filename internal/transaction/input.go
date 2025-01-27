package transaction

type CreateTransactionInput struct {
	RegisterID      string  `json:"register_id" binding:"required"`
	RegistrationFee float64 `json:"registration_fee" binding:"required"`
	ExaminationFee  float64 `json:"examination_fee" binding:"required"`
	TotalFee        float64 `json:"total_fee" binding:"required"`
	Discount        float64 `json:"discount" binding:"required"`
	TotalPayment    float64 `json:"total_payment" binding:"required"`
	PaymentType     string  `json:"payment_type" binding:"required"`
	PaymentStatus   string  `json:"payment_status" binding:"required"`
}

type UpdateTransactionInput struct {
	RegistrationFee float64 `json:"registration_fee" binding:"required"`
	ExaminationFee  float64 `json:"examination_fee" binding:"required"`
	TotalFee        float64 `json:"total_fee" binding:"required"`
	Discount        float64 `json:"discount" binding:"required"`
	TotalPayment    float64 `json:"total_payment" binding:"required"`
	PaymentType     string  `json:"payment_type" binding:"required"`
	PaymentStatus   string  `json:"payment_status" binding:"required"`
}

type GetTransactionInput struct {
	TransactionID int `uri:"transaction_id" binding:"required"`
}

type GetTransactionByMedicalRecordNoInput struct {
	MedicalRecordNo string `uri:"medical_record_no" binding:"required"`
}
