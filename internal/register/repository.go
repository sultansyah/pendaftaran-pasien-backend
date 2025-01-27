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
	query := "SELECT COUNT(register_id) AS total from register"
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
	query := "SELECT register_id, medical_record_no, session_polyclinic, clinic_id, doctor_id, department, class, entry_method, admission_fee, medical_procedure, created_at, updated_at FROM register WHERE is_deleted = 0"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Register{}, err
	}
	defer rows.Close()

	var registers []Register
	for rows.Next() {
		var register Register
		if err := rows.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt); err != nil {
			return []Register{}, err
		}

		registers = append(registers, register)
	}

	return registers, nil
}

func (r *RegisterRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, registerId string) (Register, error) {
	query := "SELECT register_id, medical_record_no, session_polyclinic, clinic_id, doctor_id, department, class, entry_method, admission_fee, medical_procedure, created_at, updated_at FROM register WHERE register_id = ? AND is_deleted = 0"
	row, err := tx.QueryContext(ctx, query, registerId)
	if err != nil {
		return Register{}, err
	}
	defer row.Close()

	var register Register
	if row.Next() {
		if err := row.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt); err != nil {
			return Register{}, err
		}

		return register, nil
	}

	return Register{}, custom.ErrNotFound
}

func (r *RegisterRepositoryImpl) FindLatestByMRNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) (Register, error) {
	query := "SELECT register_id, medical_record_no, session_polyclinic, clinic_id, doctor_id, department, class, entry_method, admission_fee, medical_procedure, created_at, updated_at FROM register WHERE medical_record_no = ? AND is_deleted = 0 ORDER BY created_at DESC LIMIT 1"
	row, err := tx.QueryContext(ctx, query, medicalRecordNo)
	if err != nil {
		return Register{}, err
	}
	defer row.Close()

	var register Register
	if row.Next() {
		if err := row.Scan(&register.RegisterID, &register.MedicalRecordNo, &register.SessionPolyclinic, &register.ClinicID, &register.DoctorID, &register.Department, &register.Class, &register.EntryMethod, &register.AdmissionFee, &register.MedicalProcedure, &register.CreatedAt, &register.UpdatedAt); err != nil {
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
