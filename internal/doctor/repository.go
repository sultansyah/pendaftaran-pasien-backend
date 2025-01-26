package doctor

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"strings"
)

type DoctorRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]Doctor, error)
	FindByDayAndClinicID(ctx context.Context, tx *sql.Tx, day string, clinicId string) ([]Doctor, error)
	FindByClinicID(ctx context.Context, tx *sql.Tx, clinicId string) ([]Doctor, error)
	FindById(ctx context.Context, tx *sql.Tx, doctorId string) (Doctor, error)
	Count(ctx context.Context, tx *sql.Tx) (int, error)
	Insert(ctx context.Context, tx *sql.Tx, doctor Doctor) (Doctor, error)
	Update(ctx context.Context, tx *sql.Tx, doctor Doctor) error
	Delete(ctx context.Context, tx *sql.Tx, doctorId string) error
}

type DoctorRepositoryImpl struct {
}

func NewDoctorRepository() DoctorRepository {
	return &DoctorRepositoryImpl{}
}

func (p *DoctorRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(doctor_id) AS total FROM doctor"
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

func (p *DoctorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, doctorId string) error {
	query := "UPDATE doctor SET is_deleted=? WHERE doctor_id = ? AND is_deleted = 0"
	_, err := tx.ExecContext(ctx, query, 1, doctorId)
	if err != nil {
		return err
	}
	return nil
}

func (p *DoctorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Doctor, error) {
	query := "SELECT doctor_id, clinic_id, doctor_name, specialization, days, start_time, end_time, phone_number, created_at, updated_at FROM doctor where is_deleted = 0"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Doctor{}, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		if err := rows.Scan(&doctor.DoctorID, &doctor.ClinicID, &doctor.DoctorName, &doctor.Specialization, &doctor.Days, &doctor.StartTime, &doctor.EndTime, &doctor.PhoneNumber, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			return []Doctor{}, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (p *DoctorRepositoryImpl) FindByDayAndClinicID(ctx context.Context, tx *sql.Tx, day string, clinicId string) ([]Doctor, error) {
	dayPattern := "%" + strings.ToLower(day) + "%"
	query := "SELECT doctor_id, clinic_id, doctor_name, specialization, days, start_time, end_time, phone_number, created_at, updated_at FROM doctor where LOWER(days) LIKE ? AND clinic_id = ? AND is_deleted = 0"
	rows, err := tx.QueryContext(ctx, query, dayPattern, clinicId)
	if err != nil {
		return []Doctor{}, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		if err := rows.Scan(&doctor.DoctorID, &doctor.ClinicID, &doctor.DoctorName, &doctor.Specialization, &doctor.Days, &doctor.StartTime, &doctor.EndTime, &doctor.PhoneNumber, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			return []Doctor{}, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (p *DoctorRepositoryImpl) FindByClinicID(ctx context.Context, tx *sql.Tx, clinicId string) ([]Doctor, error) {
	query := "SELECT doctor_id, clinic_id, doctor_name, specialization, days, start_time, end_time, phone_number, created_at, updated_at FROM doctor where clinic_id = ? AND is_deleted = 0"
	rows, err := tx.QueryContext(ctx, query, clinicId)
	if err != nil {
		return []Doctor{}, err
	}
	defer rows.Close()

	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		if err := rows.Scan(&doctor.DoctorID, &doctor.ClinicID, &doctor.DoctorName, &doctor.Specialization, &doctor.Days, &doctor.StartTime, &doctor.EndTime, &doctor.PhoneNumber, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			return []Doctor{}, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (p *DoctorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, doctorId string) (Doctor, error) {
	query := "SELECT doctor_id, clinic_id, doctor_name, specialization, days, start_time, end_time, phone_number, created_at, updated_at FROM doctor where doctor_id = ? AND is_deleted = 0"
	row, err := tx.QueryContext(ctx, query, doctorId)
	if err != nil {
		return Doctor{}, err
	}
	defer row.Close()

	var doctor Doctor
	if row.Next() {
		if err := row.Scan(&doctor.DoctorID, &doctor.ClinicID, &doctor.DoctorName, &doctor.Specialization, &doctor.Days, &doctor.StartTime, &doctor.EndTime, &doctor.PhoneNumber, &doctor.CreatedAt, &doctor.UpdatedAt); err != nil {
			return Doctor{}, err
		}
		return doctor, nil
	}

	return Doctor{}, custom.ErrNotFound
}

func (p *DoctorRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, doctor Doctor) (Doctor, error) {
	query := "INSERT INTO doctor(doctor_id, clinic_id, doctor_name, specialization, days, start_time, end_time, phone_number) VALUES (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, doctor.DoctorID, doctor.ClinicID, doctor.DoctorName, doctor.Specialization, doctor.Days, doctor.StartTime, doctor.EndTime, doctor.PhoneNumber)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (p *DoctorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, doctor Doctor) error {
	query := "UPDATE doctor SET clinic_id=?,doctor_name=?,specialization=?,days=?,start_time=?,end_time=?,phone_number=? WHERE doctor_id = ? AND is_deleted = 0"
	_, err := tx.ExecContext(ctx, query, doctor.ClinicID, doctor.DoctorName, doctor.Specialization, doctor.Days, doctor.StartTime, doctor.EndTime, doctor.PhoneNumber, doctor.DoctorID)
	if err != nil {
		return err
	}

	return nil
}
