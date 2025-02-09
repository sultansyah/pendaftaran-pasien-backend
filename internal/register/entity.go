package register

import (
	"pendaftaran-pasien-backend/internal/queue"
	"pendaftaran-pasien-backend/internal/transaction"
	"time"
)

type Register struct {
	RegisterID        string                  `json:"register_id"`
	MedicalRecordNo   string                  `json:"medical_record_no"`
	SessionPolyclinic string                  `json:"session_polyclinic"`
	ClinicID          string                  `json:"clinic_id"`
	DoctorID          string                  `json:"doctor_id"`
	Department        string                  `json:"department"`
	Class             string                  `json:"class"`
	EntryMethod       string                  `json:"entry_method"`
	AdmissionFee      string                  `json:"admission_fee"`
	MedicalProcedure  string                  `json:"medical_procedure"`
	Transaction       transaction.Transaction `json:"transaction"`
	Queue             queue.Queue             `json:"queue"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
}
