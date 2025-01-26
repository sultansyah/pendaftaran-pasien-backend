package patient

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/helper"
)

type PatientService interface {
	GetAll(ctx context.Context) ([]Patient, error)
	GetById(ctx context.Context, input GetPatientInput) (Patient, error)
	Create(ctx context.Context, input CreatePatientInput) (Patient, error)
	Update(ctx context.Context, inputId GetPatientInput, inputData CreatePatientInput) error
	Delete(ctx context.Context, input GetPatientInput) error
}

type PatientServiceImpl struct {
	DB                *sql.DB
	PatientRepository PatientRepository
}

func NewPatientService(DB *sql.DB, patientRepository PatientRepository) PatientService {
	return &PatientServiceImpl{
		DB:                DB,
		PatientRepository: patientRepository,
	}
}

func (p *PatientServiceImpl) Create(ctx context.Context, input CreatePatientInput) (Patient, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Patient{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	total, err := p.PatientRepository
}

func (p *PatientServiceImpl) Delete(ctx context.Context, input GetPatientInput) error {
	panic("unimplemented")
}

func (p *PatientServiceImpl) GetAll(ctx context.Context) ([]Patient, error) {
	panic("unimplemented")
}

func (p *PatientServiceImpl) GetById(ctx context.Context, input GetPatientInput) (Patient, error) {
	panic("unimplemented")
}

func (p *PatientServiceImpl) Update(ctx context.Context, inputId GetPatientInput, inputData CreatePatientInput) error {
	panic("unimplemented")
}
