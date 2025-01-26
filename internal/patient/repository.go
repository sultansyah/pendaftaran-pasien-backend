package patient

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type PatientRepository interface {
	FindByNoMR(ctx context.Context, tx *sql.Tx, medicalRecordNo string) (Patient, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]Patient, error)
	Insert(ctx context.Context, tx *sql.Tx, patient Patient) (Patient, error)
	Update(ctx context.Context, tx *sql.Tx, patient Patient) error
	Delete(ctx context.Context, tx *sql.Tx, medicalRecordNo string) error
	Count(ctx context.Context, tx *sql.Tx) (int, error)
}

type PatientRepositoryImpl struct {
}

func NewPatientRepository() PatientRepository {
	return &PatientRepositoryImpl{}
}

func (p *PatientRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) (int, error) {
	query := "SELECT COUNT(medical_record_no) AS total FROM patient"
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

func (p *PatientRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, medicalRecordNo string) error {
	query := "delete from patient where medical_record_no = ?"
	_, err := tx.ExecContext(ctx, query, medicalRecordNo)
	if err != nil {
		return err
	}
	return nil
}

func (p *PatientRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Patient, error) {
	query := "SELECT medical_record_no, patient_name, gender, place_of_birth, date_of_birth, address, phone_number, identity_type, identity_number, city, postal_code, medical_record_date, birth_weight, ethnicity, subdistrict, district, regency, province, citizenship, country, language, blood_type, KK_number, marital_status, religion, occupation, education, npwp, file_location, relative_name, relative_relationship, relative_phone, relative_identity_number, relative_occupation, relative_address, relative_city, relative_postal_code, mother_medical_record_no, created_at, updated_at FROM patient"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Patient{}, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var patient Patient
		if err := rows.Scan(&patient.MedicalRecordNo, &patient.PatientName, &patient.Gender, &patient.PlaceOfBirth, &patient.DateOfBirth, &patient.Address, &patient.PhoneNumber, &patient.IdentityType, &patient.IdentityNumber, &patient.City, &patient.PostalCode, &patient.MedicalRecordDate, &patient.BirthWeight, &patient.Ethnicity, &patient.Subdistrict, &patient.District, &patient.REGency, &patient.Province, &patient.Citizenship, &patient.Country, &patient.Language, &patient.BloodType, &patient.KKNumber, &patient.MaritalStatus, &patient.Religion, &patient.Occupation, &patient.Education, &patient.NPWP, &patient.FileLocation, &patient.RelativeName, &patient.RelativeRelationship, &patient.RelativePhone, &patient.RelativeIdentityNumber, &patient.RelativeOccupation, &patient.RelativeAddress, &patient.RelativeCity, &patient.RelativePostalCode, &patient.MotherMedicalRecordNo, &patient.CreatedAt, &patient.UpdatedAt); err != nil {
			return []Patient{}, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (p *PatientRepositoryImpl) FindByNoMR(ctx context.Context, tx *sql.Tx, medicalRecordNo string) (Patient, error) {
	query := "SELECT medical_record_no, patient_name, gender, place_of_birth, date_of_birth, address, phone_number, identity_type, identity_number, city, postal_code, medical_record_date, birth_weight, ethnicity, subdistrict, district, regency, province, citizenship, country, language, blood_type, KK_number, marital_status, religion, occupation, education, npwp, file_location, relative_name, relative_relationship, relative_phone, relative_identity_number, relative_occupation, relative_address, relative_city, relative_postal_code, mother_medical_record_no, created_at, updated_at FROM patient where medical_record_no = ?"
	row, err := tx.QueryContext(ctx, query, medicalRecordNo)
	if err != nil {
		return Patient{}, err
	}
	defer row.Close()

	var patient Patient
	if row.Next() {
		if err := row.Scan(&patient.MedicalRecordNo, &patient.PatientName, &patient.Gender, &patient.PlaceOfBirth, &patient.DateOfBirth, &patient.Address, &patient.PhoneNumber, &patient.IdentityType, &patient.IdentityNumber, &patient.City, &patient.PostalCode, &patient.MedicalRecordDate, &patient.BirthWeight, &patient.Ethnicity, &patient.Subdistrict, &patient.District, &patient.REGency, &patient.Province, &patient.Citizenship, &patient.Country, &patient.Language, &patient.BloodType, &patient.KKNumber, &patient.MaritalStatus, &patient.Religion, &patient.Occupation, &patient.Education, &patient.NPWP, &patient.FileLocation, &patient.RelativeName, &patient.RelativeRelationship, &patient.RelativePhone, &patient.RelativeIdentityNumber, &patient.RelativeOccupation, &patient.RelativeAddress, &patient.RelativeCity, &patient.RelativePostalCode, &patient.MotherMedicalRecordNo, &patient.CreatedAt, &patient.UpdatedAt); err != nil {
			return Patient{}, err
		}

		return patient, nil
	}

	return patient, custom.ErrNotFound
}

func (p *PatientRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, patient Patient) (Patient, error) {
	query := `INSERT INTO patient(medical_record_no, patient_name, gender, place_of_birth, date_of_birth, address, 
	phone_number, identity_type, identity_number, city, postal_code, medical_record_date, birth_weight, ethnicity, 
	subdistrict, district, regency, province, citizenship, country, language, blood_type, KK_number, marital_status, 
	religion, occupation, education, npwp, file_location, relative_name, relative_relationship, relative_phone, 
	relative_identity_number, relative_occupation, relative_address, relative_city, relative_postal_code, 
	mother_medical_record_no) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, patient.MedicalRecordNo, patient.PatientName, patient.Gender, patient.PlaceOfBirth,
		patient.DateOfBirth, patient.Address, patient.PhoneNumber, patient.IdentityType, patient.IdentityNumber, patient.City,
		patient.PostalCode, patient.MedicalRecordDate, patient.BirthWeight, patient.Ethnicity, patient.Religion,
		patient.Occupation, patient.Education, patient.NPWP, patient.FileLocation, patient.RelativeName, patient.RelativeRelationship,
		patient.RelativePhone, patient.RelativeIdentityNumber, patient.RelativeOccupation, patient.RelativeAddress,
		patient.RelativeCity, patient.RelativePostalCode, patient.MotherMedicalRecordNo)

	if err != nil {
		return patient, err

	}
	return patient, nil
}

func (p *PatientRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, patient Patient) error {
	query := `UPDATE patient SET patient_name=?,gender=?,place_of_birth=?,date_of_birth=?,
	address=?,phone_number=?,identity_type=?,identity_number=?,city=?,postal_code=?,medical_record_date=?,
	birth_weight=?,ethnicity=?,subdistrict=?,district=?,regency=?,province=?,citizenship=?,country=?,language=?,
	blood_type=?,KK_number=?,marital_status=?,religion=?,occupation=?,education=?,npwp=?,file_location=?,
	relative_name=?,relative_relationship=?,relative_phone=?,relative_identity_number=?,relative_occupation=?,
	relative_address=?,relative_city=?,relative_postal_code=?,mother_medical_record_no=? WHERE medical_record_no = ?`

	_, err := tx.ExecContext(ctx, query, patient.MedicalRecordNo, patient.PatientName, patient.Gender, patient.PlaceOfBirth,
		patient.DateOfBirth, patient.Address, patient.PhoneNumber, patient.IdentityType, patient.IdentityNumber, patient.City,
		patient.PostalCode, patient.MedicalRecordDate, patient.BirthWeight, patient.Ethnicity, patient.Religion,
		patient.Occupation, patient.Education, patient.NPWP, patient.FileLocation, patient.RelativeName, patient.RelativeRelationship,
		patient.RelativePhone, patient.RelativeIdentityNumber, patient.RelativeOccupation, patient.RelativeAddress,
		patient.RelativeCity, patient.RelativePostalCode, patient.MotherMedicalRecordNo)

	if err != nil {
		return err
	}

	return nil
}
