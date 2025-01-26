package doctor

import (
	"context"
	"database/sql"
	"fmt"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
	"pendaftaran-pasien-backend/internal/polyclinic"
)

type DoctorService interface {
	GetAll(ctx context.Context) ([]Doctor, error)
	GetById(ctx context.Context, input GetDoctorInput) (Doctor, error)
	GetByClinicId(ctx context.Context, input GetDoctorByClinicInput) ([]Doctor, error)
	GetByDayAndClinicId(ctx context.Context, input GetDoctorByDayAndClinicInput) ([]Doctor, error)
	Create(ctx context.Context, input CreateDoctorInput) (Doctor, error)
	Update(ctx context.Context, inputId GetDoctorInput, inputData CreateDoctorInput) error
	Delete(ctx context.Context, input GetDoctorInput) error
}

type DoctorServiceImpl struct {
	DB                   *sql.DB
	DoctorRepository     DoctorRepository
	PolyclinicRepository polyclinic.PolyclinicRepository
}

func NewDoctorService(DB *sql.DB, doctorRepository DoctorRepository, polyclinicRepository polyclinic.PolyclinicRepository) DoctorService {
	return &DoctorServiceImpl{
		DB:                   DB,
		DoctorRepository:     doctorRepository,
		PolyclinicRepository: polyclinicRepository,
	}
}

func (p *DoctorServiceImpl) Create(ctx context.Context, input CreateDoctorInput) (Doctor, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Doctor{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, input.ClinicID)
	if err != nil && err != custom.ErrNotFound {
		return Doctor{}, err
	}
	if err == custom.ErrNotFound || polyclinic.ClinicID == "" {
		return Doctor{}, custom.ErrPolyclinicNotFound
	}

	totalRow, err := p.DoctorRepository.Count(ctx, tx)
	if err != nil {
		return Doctor{}, err
	}

	doctor := Doctor{
		ClinicID:       input.ClinicID,
		DoctorName:     input.DoctorName,
		Specialization: input.Specialization,
		Days:           input.Days,
		StartTime:      input.StartTime,
		EndTime:        input.EndTime,
		PhoneNumber:    input.PhoneNumber,
	}

	if totalRow+1 < 10 {
		doctor.DoctorID = fmt.Sprintf("DR0%d", totalRow+1)
	} else {
		doctor.DoctorID = fmt.Sprintf("DR%d", totalRow+1)
	}

	doctor, err = p.DoctorRepository.Insert(ctx, tx, doctor)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (p *DoctorServiceImpl) Delete(ctx context.Context, input GetDoctorInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	doctor, err := p.DoctorRepository.FindById(ctx, tx, input.DoctorID)
	if err != nil {
		return err
	}

	err = p.DoctorRepository.Delete(ctx, tx, doctor.DoctorID)
	if err != nil {
		return err
	}

	return nil
}

func (p *DoctorServiceImpl) GetAll(ctx context.Context) ([]Doctor, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Doctor{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	doctors, err := p.DoctorRepository.FindAll(ctx, tx)
	if err != nil {
		return []Doctor{}, err
	}

	return doctors, nil
}

func (p *DoctorServiceImpl) GetById(ctx context.Context, input GetDoctorInput) (Doctor, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Doctor{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	doctor, err := p.DoctorRepository.FindById(ctx, tx, input.DoctorID)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (p *DoctorServiceImpl) GetByClinicId(ctx context.Context, input GetDoctorByClinicInput) ([]Doctor, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Doctor{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	doctors, err := p.DoctorRepository.FindByClinicID(ctx, tx, input.ClinicID)
	if err != nil {
		return []Doctor{}, err
	}

	return doctors, nil
}

func (p *DoctorServiceImpl) GetByDayAndClinicId(ctx context.Context, input GetDoctorByDayAndClinicInput) ([]Doctor, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Doctor{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, input.ClinicID)
	if err != nil && err != custom.ErrNotFound {
		return []Doctor{}, err
	}
	if err == custom.ErrNotFound || polyclinic.ClinicID == "" {
		return []Doctor{}, custom.ErrPolyclinicNotFound
	}

	doctors, err := p.DoctorRepository.FindByDayAndClinicID(ctx, tx, input.Day, input.ClinicID)
	if err != nil {
		return []Doctor{}, err
	}

	return doctors, nil
}

func (p *DoctorServiceImpl) Update(ctx context.Context, inputId GetDoctorInput, inputData CreateDoctorInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, inputData.ClinicID)
	if err != nil && err != custom.ErrNotFound {
		return err
	}
	if err == custom.ErrNotFound || polyclinic.ClinicID == "" {
		return custom.ErrPolyclinicNotFound
	}

	doctor, err := p.DoctorRepository.FindById(ctx, tx, inputId.DoctorID)
	if err != nil {
		return err
	}

	doctor.ClinicID = inputData.ClinicID
	doctor.DoctorName = inputData.DoctorName
	doctor.Specialization = inputData.Specialization
	doctor.Days = inputData.Days
	doctor.StartTime = inputData.StartTime
	doctor.EndTime = inputData.EndTime
	doctor.PhoneNumber = inputData.PhoneNumber

	err = p.DoctorRepository.Update(ctx, tx, doctor)
	if err != nil {
		return err
	}

	return nil
}
