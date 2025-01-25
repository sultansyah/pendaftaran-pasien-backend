package polyclinic

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type PolyclinicRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]Polyclinic, error)
	FindById(ctx context.Context, tx *sql.Tx, clinic_id string) (Polyclinic, error)
	Count(ctx context.Context, tx *sql.Tx) (int, error)
	Insert(ctx context.Context, tx *sql.Tx, polyclinic Polyclinic) (Polyclinic, error)
	Update(ctx context.Context, tx *sql.Tx, polyclinic Polyclinic) error
	Delete(ctx context.Context, tx *sql.Tx, clinic_id string) error
}

type PolyclinicRepositoryImpl struct {
}

func NewPolyclinicRepository() PolyclinicRepository {
	return &PolyclinicRepositoryImpl{}
}

func (p *PolyclinicRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(clinic_id) AS total FROM polyclinic"
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

func (p *PolyclinicRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, clinic_id string) error {
	query := "UPDATE polyclinic SET is_deleted=? WHERE clinic_id = ? AND is_deleted = 0"
	_, err := tx.ExecContext(ctx, query, 1, clinic_id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PolyclinicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Polyclinic, error) {
	query := "SELECT clinic_id, clinic_name, location, created_at, updated_at FROM polyclinic where is_deleted = 0"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Polyclinic{}, err
	}
	defer rows.Close()

	var polyclinics []Polyclinic
	for rows.Next() {
		var polyclinic Polyclinic
		if err := rows.Scan(&polyclinic.ClinicID, &polyclinic.ClinicName, &polyclinic.Location, &polyclinic.CreatedAt, &polyclinic.UpdatedAt); err != nil {
			return []Polyclinic{}, err
		}
		polyclinics = append(polyclinics, polyclinic)
	}

	return polyclinics, nil
}

func (p *PolyclinicRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, clinic_id string) (Polyclinic, error) {
	query := "SELECT clinic_id, clinic_name, location, created_at, updated_at FROM polyclinic where clinic_id = ? AND is_deleted = 0"
	row, err := tx.QueryContext(ctx, query, clinic_id)
	if err != nil {
		return Polyclinic{}, err
	}
	defer row.Close()

	var polyclinic Polyclinic
	if row.Next() {
		if err := row.Scan(&polyclinic.ClinicID, &polyclinic.ClinicName, &polyclinic.Location, &polyclinic.CreatedAt, &polyclinic.UpdatedAt); err != nil {
			return Polyclinic{}, err
		}
		return polyclinic, nil
	}

	return Polyclinic{}, custom.ErrNotFound
}

func (p *PolyclinicRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, polyclinic Polyclinic) (Polyclinic, error) {
	query := "INSERT INTO polyclinic(clinic_id, clinic_name, location) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, query, polyclinic.ClinicID, polyclinic.ClinicName, polyclinic.Location)
	if err != nil {
		return polyclinic, err
	}

	return polyclinic, nil
}

func (p *PolyclinicRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, polyclinic Polyclinic) error {
	query := "UPDATE polyclinic SET clinic_name=?,location=? WHERE clinic_id = ? AND is_deleted = 0"
	_, err := tx.ExecContext(ctx, query, polyclinic.ClinicName, polyclinic.Location, polyclinic.ClinicID)
	if err != nil {
		return err
	}

	return nil
}
