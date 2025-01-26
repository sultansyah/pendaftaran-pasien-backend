package patient

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
)

type PatientService interface {
	GetAll(ctx context.Context) ([]Patient, error)
	GetByNoMR(ctx context.Context, input GetPatientInput) (Patient, error)
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

	total, err := p.PatientRepository.Count(ctx, tx)
	if err != nil {
		return Patient{}, err
	}
	dateOfBirth, err := helper.ParseToHour(input.DateOfBirth)
	if err != nil {
		return Patient{}, err
	}

	medicalRecordDate, err := helper.ParseDateTimeLocal(input.MedicalRecordDate)
	if err != nil {
		return Patient{}, err
	}

	medicalRecordNo := helper.GenerateMedicalRecordNo(total)

	patient := Patient{
		MedicalRecordNo:        medicalRecordNo,
		PatientName:            input.PatientName,
		Gender:                 input.Gender,
		PlaceOfBirth:           input.PlaceOfBirth,
		DateOfBirth:            dateOfBirth,
		Address:                input.Address,
		PhoneNumber:            input.PhoneNumber,
		IdentityType:           input.IdentityType,
		IdentityNumber:         input.IdentityNumber,
		City:                   input.City,
		PostalCode:             input.PostalCode,
		MedicalRecordDate:      medicalRecordDate,
		BirthWeight:            input.BirthWeight,
		Ethnicity:              input.Ethnicity,
		Subdistrict:            input.Subdistrict,
		District:               input.District,
		REGency:                input.REGency,
		Province:               input.Province,
		Citizenship:            input.Citizenship,
		Country:                input.Country,
		Language:               input.Language,
		BloodType:              input.BloodType,
		KKNumber:               input.KKNumber,
		MaritalStatus:          input.MaritalStatus,
		Religion:               input.Religion,
		Occupation:             input.Occupation,
		Education:              input.Education,
		NPWP:                   input.NPWP,
		FileLocation:           input.FileLocation,
		RelativeName:           input.RelativeName,
		RelativeRelationship:   input.RelativeRelationship,
		RelativePhone:          input.RelativePhone,
		RelativeIdentityNumber: input.RelativeIdentityNumber,
		RelativeOccupation:     input.RelativeOccupation,
		RelativeAddress:        input.RelativeAddress,
		RelativeCity:           input.RelativeCity,
		RelativePostalCode:     input.RelativePostalCode,
		MotherMedicalRecordNo:  input.MotherMedicalRecordNo,
	}

	patient, err = p.PatientRepository.Insert(ctx, tx, patient)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

func (p *PatientServiceImpl) Delete(ctx context.Context, input GetPatientInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	patient, err := p.PatientRepository.FindByNoMR(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return err
	}
	if patient.MedicalRecordNo == "" {
		return custom.ErrNotFound
	}

	err = p.PatientRepository.Delete(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return err
	}

	return nil
}

func (p *PatientServiceImpl) GetAll(ctx context.Context) ([]Patient, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Patient{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	patients, err := p.PatientRepository.FindAll(ctx, tx)
	if err != nil {
		return []Patient{}, err
	}

	return patients, nil
}

func (p *PatientServiceImpl) GetByNoMR(ctx context.Context, input GetPatientInput) (Patient, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return Patient{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	patient, err := p.PatientRepository.FindByNoMR(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

func (p *PatientServiceImpl) Update(ctx context.Context, inputId GetPatientInput, inputData CreatePatientInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	patient, err := p.PatientRepository.FindByNoMR(ctx, tx, inputId.MedicalRecordNo)
	if err != nil {
		return err
	}
	if patient.MedicalRecordNo == "" {
		return custom.ErrNotFound
	}

	dateOfBirth, err := helper.ParseToHour(inputData.DateOfBirth)
	if err != nil {
		return err
	}

	medicalRecordDate, err := helper.ParseDateTimeLocal(inputData.MedicalRecordDate)
	if err != nil {
		return err
	}

	patient = Patient{
		PatientName:            inputData.PatientName,
		Gender:                 inputData.Gender,
		PlaceOfBirth:           inputData.PlaceOfBirth,
		DateOfBirth:            dateOfBirth,
		Address:                inputData.Address,
		PhoneNumber:            inputData.PhoneNumber,
		IdentityType:           inputData.IdentityType,
		IdentityNumber:         inputData.IdentityNumber,
		City:                   inputData.City,
		PostalCode:             inputData.PostalCode,
		MedicalRecordDate:      medicalRecordDate,
		BirthWeight:            inputData.BirthWeight,
		Ethnicity:              inputData.Ethnicity,
		Subdistrict:            inputData.Subdistrict,
		District:               inputData.District,
		REGency:                inputData.REGency,
		Province:               inputData.Province,
		Citizenship:            inputData.Citizenship,
		Country:                inputData.Country,
		Language:               inputData.Language,
		BloodType:              inputData.BloodType,
		KKNumber:               inputData.KKNumber,
		MaritalStatus:          inputData.MaritalStatus,
		Religion:               inputData.Religion,
		Occupation:             inputData.Occupation,
		Education:              inputData.Education,
		NPWP:                   inputData.NPWP,
		FileLocation:           inputData.FileLocation,
		RelativeName:           inputData.RelativeName,
		RelativeRelationship:   inputData.RelativeRelationship,
		RelativePhone:          inputData.RelativePhone,
		RelativeIdentityNumber: inputData.RelativeIdentityNumber,
		RelativeOccupation:     inputData.RelativeOccupation,
		RelativeAddress:        inputData.RelativeAddress,
		RelativeCity:           inputData.RelativeCity,
		RelativePostalCode:     inputData.RelativePostalCode,
		MotherMedicalRecordNo:  inputData.MotherMedicalRecordNo,
	}

	err = p.PatientRepository.Update(ctx, tx, patient)
	if err != nil {
		return err
	}

	return nil
}
