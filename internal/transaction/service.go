package transaction

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/helper"
)

type TransactionService interface {
	GetAll(ctx context.Context) ([]Transaction, error)
	GetById(ctx context.Context, input GetTransactionInput) (Transaction, error)
	GetByMedicalRecordNo(ctx context.Context, input GetTransactionByMedicalRecordNoInput) ([]Transaction, error)
	Create(ctx context.Context, input CreateTransactionInput) (Transaction, error)
	Update(ctx context.Context, inputID GetTransactionInput, inputData UpdateTransactionInput) error
}

type TransactionServiceImpl struct {
	DB                    *sql.DB
	TransactionRepository TransactionRepository
}

func NewTransactionService(DB *sql.DB, transactionRepository TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		DB:                    DB,
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionServiceImpl) Create(ctx context.Context, input CreateTransactionInput) (Transaction, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return Transaction{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	transaction := Transaction{
		RegisterID:      input.RegisterID,
		RegistrationFee: input.RegistrationFee,
		ExaminationFee:  input.ExaminationFee,
		TotalFee:        input.TotalFee,
		Discount:        input.Discount,
		TotalPayment:    input.TotalPayment,
		PaymentType:     input.PaymentType,
		PaymentStatus:   input.PaymentStatus,
	}

	transaction, err = t.TransactionRepository.Insert(ctx, tx, transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (t *TransactionServiceImpl) GetAll(ctx context.Context) ([]Transaction, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Transaction{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	transactions, err := t.TransactionRepository.FindAll(ctx, tx)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}

func (t *TransactionServiceImpl) GetById(ctx context.Context, input GetTransactionInput) (Transaction, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return Transaction{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	transaction, err := t.TransactionRepository.FindById(ctx, tx, input.TransactionID)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (t *TransactionServiceImpl) GetByMedicalRecordNo(ctx context.Context, input GetTransactionByMedicalRecordNoInput) ([]Transaction, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Transaction{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	transactions, err := t.TransactionRepository.FindByMedicalRecordNo(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}

func (t *TransactionServiceImpl) Update(ctx context.Context, inputID GetTransactionInput, inputData UpdateTransactionInput) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	transaction, err := t.TransactionRepository.FindById(ctx, tx, inputID.TransactionID)
	if err != nil {
		return err
	}

	transaction.RegistrationFee = inputData.RegistrationFee
	transaction.ExaminationFee = inputData.ExaminationFee
	transaction.TotalFee = inputData.TotalFee
	transaction.Discount = inputData.Discount
	transaction.TotalPayment = inputData.TotalPayment
	transaction.PaymentType = inputData.PaymentType
	transaction.PaymentStatus = inputData.PaymentStatus

	err = t.TransactionRepository.Update(ctx, tx, transaction)
	if err != nil {
		return err
	}

	return nil
}
