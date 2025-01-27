package register

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/doctor"
	"pendaftaran-pasien-backend/internal/helper"
	"pendaftaran-pasien-backend/internal/patient"
	"pendaftaran-pasien-backend/internal/polyclinic"
	"pendaftaran-pasien-backend/internal/queue"
	"pendaftaran-pasien-backend/internal/transaction"
)

type RegisterService interface {
	GetAll(ctx context.Context) ([]Register, error)
	GetById(ctx context.Context, input GetRegisterInput) (Register, error)
	GetLatestByMRNo(ctx context.Context, input GetRegisterByMRNoInput) (Register, error)
	Create(ctx context.Context, input CreateRegisterInput) (RegisterWithTransactionAndQueue, error)
	Update(ctx context.Context, inputId GetRegisterInput, inputData UpdateRegisterInput) (RegisterWithTransactionAndQueue, error)
	Delete(ctx context.Context, input GetRegisterInput) error
}

type RegisterServiceImpl struct {
	DB                    *sql.DB
	RegisterRepository    RegisterRepository
	PolyclinicRepository  polyclinic.PolyclinicRepository
	DoctorRepository      doctor.DoctorRepository
	PatientRepository     patient.PatientRepository
	QueueRepository       queue.QueueRepository
	TransactionRepository transaction.TransactionRepository
}

func NewRegisterService(DB *sql.DB, registerRepository RegisterRepository, polyclinicRepository polyclinic.PolyclinicRepository, doctorRepository doctor.DoctorRepository, patientRepository patient.PatientRepository, queueRepository queue.QueueRepository, transactionRepository transaction.TransactionRepository) RegisterService {
	return &RegisterServiceImpl{
		DB:                    DB,
		RegisterRepository:    registerRepository,
		PolyclinicRepository:  polyclinicRepository,
		DoctorRepository:      doctorRepository,
		PatientRepository:     patientRepository,
		QueueRepository:       queueRepository,
		TransactionRepository: transactionRepository,
	}
}

func (r *RegisterServiceImpl) Create(ctx context.Context, input CreateRegisterInput) (RegisterWithTransactionAndQueue, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	patient, err := r.PatientRepository.FindByNoMR(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if patient.MedicalRecordNo == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrMedicalRecordNotFound
	}

	polyclinic, err := r.PolyclinicRepository.FindById(ctx, tx, input.ClinicID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if polyclinic.ClinicID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrPolyclinicNotFound
	}

	doctor, err := r.DoctorRepository.FindById(ctx, tx, input.DoctorID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if doctor.DoctorID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrDoctorNotFound
	}

	total, err := r.RegisterRepository.Count(ctx, tx)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	register := Register{
		RegisterID:        helper.GenerateRegisterNo(total + 1),
		MedicalRecordNo:   input.MedicalRecordNo,
		SessionPolyclinic: input.SessionPolyclinic,
		ClinicID:          polyclinic.ClinicID,
		DoctorID:          doctor.DoctorID,
		Department:        input.Department,
		Class:             input.Class,
		EntryMethod:       input.EntryMethod,
		AdmissionFee:      input.AdmissionFee,
		MedicalProcedure:  input.MedicalProcedure,
	}

	register, err = r.RegisterRepository.Insert(ctx, tx, register)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	register, err = r.RegisterRepository.FindById(ctx, tx, register.RegisterID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if register.RegisterID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrNotFound
	}

	date, err := helper.ParseDatetimeToDate(register.CreatedAt)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	total, err = r.QueueRepository.CountQueueByDay(ctx, tx, date)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	queue := queue.Queue{
		RegisterID:  register.RegisterID,
		QueueNumber: total + 1,
	}

	queue, err = r.QueueRepository.Insert(ctx, tx, queue)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	transaction := transaction.Transaction{
		RegisterID:      register.RegisterID,
		RegistrationFee: input.RegistrationFee,
		ExaminationFee:  input.ExaminationFee,
		TotalFee:        input.TotalFee,
		Discount:        input.Discount,
		TotalPayment:    input.TotalPayment,
		PaymentType:     input.PaymentType,
		PaymentStatus:   input.PaymentStatus,
	}

	transaction, err = r.TransactionRepository.Insert(ctx, tx, transaction)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	return RegisterWithTransactionAndQueueFormatter(register, queue, transaction, patient), nil
}

func (r *RegisterServiceImpl) Delete(ctx context.Context, input GetRegisterInput) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	register, err := r.RegisterRepository.FindById(ctx, tx, input.RegisterID)
	if err != nil {
		return err
	}

	err = r.RegisterRepository.Delete(ctx, tx, register.RegisterID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RegisterServiceImpl) GetAll(ctx context.Context) ([]Register, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Register{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	registers, err := r.RegisterRepository.FindAll(ctx, tx)
	if err != nil {
		return []Register{}, err
	}

	return registers, nil
}

func (r *RegisterServiceImpl) GetById(ctx context.Context, input GetRegisterInput) (Register, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return Register{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	register, err := r.RegisterRepository.FindById(ctx, tx, input.RegisterID)
	if err != nil {
		return Register{}, err
	}

	return register, nil
}

func (r *RegisterServiceImpl) GetLatestByMRNo(ctx context.Context, input GetRegisterByMRNoInput) (Register, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return Register{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	patient, err := r.PatientRepository.FindByNoMR(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return Register{}, err
	}

	register, err := r.RegisterRepository.FindLatestByMRNo(ctx, tx, patient.MedicalRecordNo)
	if err != nil {
		return Register{}, err
	}

	return register, nil
}

func (r *RegisterServiceImpl) Update(ctx context.Context, inputId GetRegisterInput, inputData UpdateRegisterInput) (RegisterWithTransactionAndQueue, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	register, err := r.RegisterRepository.FindById(ctx, tx, inputId.RegisterID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if register.RegisterID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrRegisterNotFound
	}

	patient, err := r.PatientRepository.FindByNoMR(ctx, tx, inputData.MedicalRecordNo)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if patient.MedicalRecordNo == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrMedicalRecordNotFound
	}

	polyclinic, err := r.PolyclinicRepository.FindById(ctx, tx, inputData.ClinicID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if polyclinic.ClinicID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrPolyclinicNotFound
	}

	doctor, err := r.DoctorRepository.FindById(ctx, tx, inputData.DoctorID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if doctor.DoctorID == "" {
		return RegisterWithTransactionAndQueue{}, custom.ErrDoctorNotFound
	}

	register.MedicalRecordNo = inputData.MedicalRecordNo
	register.SessionPolyclinic = inputData.SessionPolyclinic
	register.ClinicID = polyclinic.ClinicID
	register.DoctorID = doctor.DoctorID
	register.Department = inputData.Department
	register.Class = inputData.Class
	register.EntryMethod = inputData.EntryMethod
	register.AdmissionFee = inputData.AdmissionFee
	register.MedicalProcedure = inputData.MedicalProcedure

	err = r.RegisterRepository.Update(ctx, tx, register)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	transaction, err := r.TransactionRepository.FindById(ctx, tx, inputData.TransactionID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if transaction.TransactionID <= 0 {
		return RegisterWithTransactionAndQueue{}, err
	}
	if transaction.RegisterID != register.RegisterID {
		return RegisterWithTransactionAndQueue{}, custom.ErrUnauthorized
	}

	transaction.RegisterID = register.RegisterID
	transaction.RegistrationFee = inputData.RegistrationFee
	transaction.ExaminationFee = inputData.ExaminationFee
	transaction.TotalFee = inputData.TotalFee
	transaction.Discount = inputData.Discount
	transaction.TotalPayment = inputData.TotalPayment
	transaction.PaymentType = inputData.PaymentType
	transaction.PaymentStatus = inputData.PaymentStatus

	err = r.TransactionRepository.Update(ctx, tx, transaction)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}

	queue, err := r.QueueRepository.FindById(ctx, tx, inputData.QueueID)
	if err != nil {
		return RegisterWithTransactionAndQueue{}, err
	}
	if queue.QueueID <= 0 {
		return RegisterWithTransactionAndQueue{}, err
	}
	if queue.RegisterID != register.RegisterID {
		return RegisterWithTransactionAndQueue{}, custom.ErrUnauthorized
	}

	return RegisterWithTransactionAndQueueFormatter(register, queue, transaction, patient), nil
}
