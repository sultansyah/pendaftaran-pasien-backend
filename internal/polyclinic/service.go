package polyclinic

import (
	"context"
	"database/sql"
	"fmt"
	"pendaftaran-pasien-backend/internal/helper"
)

type PolyclinicService interface {
	GetAll(ctx context.Context) ([]Polyclinic, error)
	GetById(ctx context.Context, input GetPolyclinicInput) (Polyclinic, error)
	Create(ctx context.Context, input CreatePolyclinicInput) (Polyclinic, error)
	Update(ctx context.Context, inputId GetPolyclinicInput, inputData CreatePolyclinicInput) error
	Delete(ctx context.Context, input GetPolyclinicInput) error
}

type PolyclinicServiceImpl struct {
	DB                   *sql.DB
	PolyclinicRepository PolyclinicRepository
}

func NewPolyclinicService(DB *sql.DB, polyclinicRepository PolyclinicRepository) PolyclinicService {
	return &PolyclinicServiceImpl{
		DB:                   DB,
		PolyclinicRepository: polyclinicRepository,
	}
}

func (p *PolyclinicServiceImpl) Create(ctx context.Context, input CreatePolyclinicInput) (Polyclinic, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Polyclinic{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	totalRow, err := p.PolyclinicRepository.Count(ctx, tx)
	if err != nil {
		return Polyclinic{}, err
	}

	polyclinic := Polyclinic{
		ClinicName: input.ClinicName,
		Location:   input.Location,
	}

	if totalRow+1 < 10 {
		polyclinic.ClinicID = fmt.Sprintf("POL0%d", totalRow+1)
	} else {
		polyclinic.ClinicID = fmt.Sprintf("POL%d", totalRow+1)
	}

	polyclinic, err = p.PolyclinicRepository.Insert(ctx, tx, polyclinic)
	if err != nil {
		return Polyclinic{}, err
	}

	return polyclinic, nil
}

func (p *PolyclinicServiceImpl) Delete(ctx context.Context, input GetPolyclinicInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, input.ClinicID)
	if err != nil {
		return err
	}

	err = p.PolyclinicRepository.Delete(ctx, tx, polyclinic.ClinicID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PolyclinicServiceImpl) GetAll(ctx context.Context) ([]Polyclinic, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Polyclinic{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinics, err := p.PolyclinicRepository.FindAll(ctx, tx)
	if err != nil {
		return []Polyclinic{}, err
	}

	return polyclinics, nil
}

func (p *PolyclinicServiceImpl) GetById(ctx context.Context, input GetPolyclinicInput) (Polyclinic, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Polyclinic{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, input.ClinicID)
	if err != nil {
		return Polyclinic{}, err
	}

	return polyclinic, nil
}

func (p *PolyclinicServiceImpl) Update(ctx context.Context, inputId GetPolyclinicInput, inputData CreatePolyclinicInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	polyclinic, err := p.PolyclinicRepository.FindById(ctx, tx, inputId.ClinicID)
	if err != nil {
		return err
	}

	polyclinic.ClinicName = inputData.ClinicName
	polyclinic.Location = inputData.Location

	err = p.PolyclinicRepository.Update(ctx, tx, polyclinic)
	if err != nil {
		return err
	}

	return nil
}
