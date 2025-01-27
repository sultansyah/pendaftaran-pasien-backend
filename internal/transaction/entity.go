package transaction

import "time"

type Transaction struct {
	TransactionID   int       `json:"transaction_id"`
	RegisterID      string    `json:"register_id"`
	RegistrationFee float64   `json:"registration_fee"`
	ExaminationFee  float64   `json:"examination_fee"`
	TotalFee        float64   `json:"total_fee"`
	Discount        float64   `json:"discount"`
	TotalPayment    float64   `json:"total_payment"`
	PaymentType     string    `json:"payment_type"`
	PaymentStatus   string    `json:"payment_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
