package register

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type RegisterRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]Register, error)
	FindById(ctx context.Context, tx *sql.Tx, registerId string) (Register, error)
	FindLatestByMRNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) (Register, error)
	Insert(ctx context.Context, tx *sql.Tx, register Register) (Register, error)
	Update(ctx context.Context, tx *sql.Tx, register Register) error
	Delete(ctx context.Context, tx *sql.Tx, registerId string) error
	Count(ctx context.Context, tx *sql.Tx) (int, error)
}

type RegisterRepositoryImpl struct {
}

func NewRegisterRepository() RegisterRepository {
	return &RegisterRepositoryImpl{}
}

func (r *RegisterRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) (int, error) {
	query := `SELECT COALESCE(MAX(CAST(REPLACE(register_id, 'RG', '') AS UNSIGNED)), 1) AS total 
FROM register`
	row, err := tx.QueryContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer row.Close()

	var total int
	if row.Next() {
		if err := row.Scan(&total); err != nil {
			return -1, err
		}

		return total, nil
	}

	return -1, custom.ErrNotFound
}

func (r *RegisterRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, registerId string) error {
	query := "UPDATE register SET is_deleted=? WHERE register_id = ? AND is_deleted = 0"
	_, err := tx.ExecContext(ctx, query, 1, registerId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RegisterRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Register, error) {
	query := `SELECT RG.register_id, RG.medical_record_no, RG.session_polyclinic, RG.clinic_id, 
RG.doctor_id, RG.department, RG.class, RG.entry_method, RG.admission_fee, RG.medical_procedure, 
RG.created_at, RG.updated_at, TR.transaction_id, TR.registration_fee, TR.examination_fee, 
TR.total_fee, TR.discount, TR.total_payment, TR.payment_type, TR.payment_status,
Q.queue_id, Q.queue_number, Q.is_completed
FROM register AS RG
INNER JOIN transactions AS TR ON RG.register_id = TR.register_id
INNER JOIN queue AS Q ON RG.register_id = Q.register_id
WHERE RG.is_deleted = 0`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Register{}, err
	}
	defer rows.Close()

	var registers []Register
	for rows.Next() {
		var register Register
		if err := rows.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt, &register.Transaction.TransactionID, &register.Transaction.RegistrationFee, &register.Transaction.ExaminationFee, &register.Transaction.TotalFee, &register.Transaction.Discount, &register.Transaction.TotalPayment, &register.Transaction.PaymentType, &register.Transaction.PaymentStatus, &register.Queue.QueueID, &register.Queue.QueueNumber, &register.Queue.IsCompleted); err != nil {
			return []Register{}, err
		}

		registers = append(registers, register)
	}

	return registers, nil
}

func (r *RegisterRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, registerId string) (Register, error) {
	query := `SELECT RG.register_id, RG.medical_record_no, RG.session_polyclinic, RG.clinic_id, 
RG.doctor_id, RG.department, RG.class, RG.entry_method, RG.admission_fee, RG.medical_procedure, 
RG.created_at, RG.updated_at, TR.transaction_id, TR.registration_fee, TR.examination_fee, 
TR.total_fee, TR.discount, TR.total_payment, TR.payment_type, TR.payment_status,
Q.queue_id, Q.queue_number, Q.is_completed
FROM register AS RG
INNER JOIN transactions AS TR ON RG.register_id = TR.register_id
INNER JOIN queue AS Q ON RG.register_id = Q.register_id
WHERE RG.register_id = ? AND RG.is_deleted = 0`
	row, err := tx.QueryContext(ctx, query, registerId)
	if err != nil {
		return Register{}, err
	}
	defer row.Close()

	var register Register
	if row.Next() {
		if err := row.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt, &register.Transaction.TransactionID, &register.Transaction.RegistrationFee, &register.Transaction.ExaminationFee, &register.Transaction.TotalFee, &register.Transaction.Discount, &register.Transaction.TotalPayment, &register.Transaction.PaymentType, &register.Transaction.PaymentStatus, &register.Queue.QueueID, &register.Queue.QueueNumber, &register.Queue.IsCompleted); err != nil {
			return Register{}, err
		}

		return register, nil
	}

	return Register{}, custom.ErrRegisterNotFound
}

func (r *RegisterRepositoryImpl) FindLatestByMRNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) (Register, error) {
	query := `SSELECT RG.register_id, RG.medical_record_no, RG.session_polyclinic, RG.clinic_id, 
RG.doctor_id, RG.department, RG.class, RG.entry_method, RG.admission_fee, RG.medical_procedure, 
RG.created_at, RG.updated_at, TR.transaction_id, TR.registration_fee, TR.examination_fee, 
TR.total_fee, TR.discount, TR.total_payment, TR.payment_type, TR.payment_status,
Q.queue_id, Q.queue_number, Q.is_completed
FROM register AS RG
INNER JOIN transactions AS TR ON RG.register_id = TR.register_id
INNER JOIN queue AS Q ON RG.register_id = Q.register_id
WHERE RG.medical_record_no = ? AND RG.is_deleted = 0 ORDER BY RG.created_at DESC LIMIT 1`
	row, err := tx.QueryContext(ctx, query, medicalRecordNo)
	if err != nil {
		return Register{}, err
	}
	defer row.Close()

	var register Register
	if row.Next() {
		if err := row.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt, &register.Transaction.TransactionID, &register.Transaction.RegistrationFee, &register.Transaction.ExaminationFee, &register.Transaction.TotalFee, &register.Transaction.Discount, &register.Transaction.TotalPayment, &register.Transaction.PaymentType, &register.Transaction.PaymentStatus); err != nil {
			return Register{}, err
		}

		return register, nil
	}

	return Register{}, custom.ErrNotFound
}

func (r *RegisterRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, register Register) (Register, error) {
	query := `INSERT INTO register(register_id, medical_record_no, session_polyclinic,
	clinic_id, doctor_id, department, class, entry_method, admission_fee, medical_procedure) 
	VALUES (?,?,?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, register.RegisterID, register.MedicalRecordNo, register.SessionPolyclinic, register.ClinicID, register.DoctorID, register.Department, register.Class, register.EntryMethod, register.AdmissionFee, register.MedicalProcedure)
	if err != nil {
		return register, err
	}

	return register, nil
}

func (r *RegisterRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, register Register) error {
	query := `UPDATE register SET medical_record_no=?,session_polyclinic=?,clinic_id=?,doctor_id=?,department=?,class=?,entry_method=?,admission_fee=?,medical_procedure=? WHERE register_id = ? AND is_deleted = 0`

	_, err := tx.ExecContext(ctx, query, register.MedicalRecordNo, register.SessionPolyclinic, register.ClinicID, register.DoctorID, register.Department, register.Class, register.EntryMethod, register.AdmissionFee, register.MedicalProcedure, register.RegisterID)
	if err != nil {
		return err
	}

	return nil
}
