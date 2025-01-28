package patient

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
)

type PatientService interface {
	GetAll(ctx context.Context, input GetPatientInput) ([]Patient, error)
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

	var motherMedicalRecordNo *string
	if input.MotherMedicalRecordNo != nil {
		motherP, err := p.PatientRepository.FindByNoMR(ctx, tx, *input.MotherMedicalRecordNo)
		if err != nil {
			return Patient{}, err
		}
		if motherP.MedicalRecordNo == "" {
			return Patient{}, custom.ErrMedicalRecordNotFound
		}

		motherMedicalRecordNo = &motherP.MedicalRecordNo
	}

	identityNumber, err := p.PatientRepository.FindByIdentityNumber(ctx, tx, input.IdentityNumber)
	if err != nil {
		return Patient{}, err
	}
	if identityNumber.MedicalRecordNo != "" && err != custom.ErrIdentityNumberNotFound {
		return Patient{}, custom.ErrIdentityNumberAlreadyExists
	}

	total, err := p.PatientRepository.Count(ctx, tx)
	if err != nil {
		return Patient{}, err
	}

	dateOfBirth, err := helper.ParseDate(input.DateOfBirth)
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
		MotherMedicalRecordNo:  motherMedicalRecordNo,
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
		return custom.ErrMedicalRecordNotFound
	}

	err = p.PatientRepository.Delete(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return err
	}

	return nil
}

func (p *PatientServiceImpl) GetAll(ctx context.Context, input GetPatientInput) ([]Patient, error) {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Patient{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	filters := helper.GenerateFilter(input)

	if medicalRecordNo, ok := filters["medical_record_no"]; ok {
		medicalRecordNo, ok := medicalRecordNo.(string)
		if !ok {
			return []Patient{}, custom.ErrInternal
		}

		_, err := p.PatientRepository.FindByNoMR(ctx, tx, medicalRecordNo)
		if err != nil {
			return []Patient{}, err
		}
	}

	patients, err := p.PatientRepository.Find(ctx, tx, filters)
	if err != nil {
		return []Patient{}, err
	}

	return patients, nil
}

func (p *PatientServiceImpl) Update(ctx context.Context, inputId GetPatientInput, inputData CreatePatientInput) error {
	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	var motherMedicalRecordNo *string
	if inputData.MotherMedicalRecordNo != nil {
		motherP, err := p.PatientRepository.FindByNoMR(ctx, tx, *inputData.MotherMedicalRecordNo)
		if err != nil {
			return err
		}
		if motherP.MedicalRecordNo == "" {
			return custom.ErrMedicalRecordNotFound
		}

		motherMedicalRecordNo = &motherP.MedicalRecordNo
	}

	patient, err := p.PatientRepository.FindByNoMR(ctx, tx, inputId.MedicalRecordNo)
	if err != nil {
		return err
	}
	if patient.MedicalRecordNo == "" {
		return custom.ErrMedicalRecordNotFound
	}

	dateOfBirth, err := helper.ParseDate(inputData.DateOfBirth)
	if err != nil {
		return err
	}

	medicalRecordDate, err := helper.ParseDateTimeLocal(inputData.MedicalRecordDate)
	if err != nil {
		return err
	}

	patient.PatientName = inputData.PatientName
	patient.Gender = inputData.Gender
	patient.PlaceOfBirth = inputData.PlaceOfBirth
	patient.DateOfBirth = dateOfBirth
	patient.Address = inputData.Address
	patient.PhoneNumber = inputData.PhoneNumber
	patient.IdentityType = inputData.IdentityType
	patient.IdentityNumber = inputData.IdentityNumber
	patient.City = inputData.City
	patient.PostalCode = inputData.PostalCode
	patient.MedicalRecordDate = medicalRecordDate
	patient.BirthWeight = inputData.BirthWeight
	patient.Ethnicity = inputData.Ethnicity
	patient.Subdistrict = inputData.Subdistrict
	patient.District = inputData.District
	patient.REGency = inputData.REGency
	patient.Province = inputData.Province
	patient.Citizenship = inputData.Citizenship
	patient.Country = inputData.Country
	patient.Language = inputData.Language
	patient.BloodType = inputData.BloodType
	patient.KKNumber = inputData.KKNumber
	patient.MaritalStatus = inputData.MaritalStatus
	patient.Religion = inputData.Religion
	patient.Occupation = inputData.Occupation
	patient.Education = inputData.Education
	patient.NPWP = inputData.NPWP
	patient.FileLocation = inputData.FileLocation
	patient.RelativeName = inputData.RelativeName
	patient.RelativeRelationship = inputData.RelativeRelationship
	patient.RelativePhone = inputData.RelativePhone
	patient.RelativeIdentityNumber = inputData.RelativeIdentityNumber
	patient.RelativeOccupation = inputData.RelativeOccupation
	patient.RelativeAddress = inputData.RelativeAddress
	patient.RelativeCity = inputData.RelativeCity
	patient.RelativePostalCode = inputData.RelativePostalCode

	if motherMedicalRecordNo != nil && motherMedicalRecordNo != patient.MotherMedicalRecordNo {
		patient.MotherMedicalRecordNo = motherMedicalRecordNo
	}

	err = p.PatientRepository.Update(ctx, tx, patient)
	if err != nil {
		return err
	}

	return nil
}
