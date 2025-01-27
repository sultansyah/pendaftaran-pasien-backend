package transaction

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type TransactionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]Transaction, error)
	FindById(ctx context.Context, tx *sql.Tx, transactionId int) (Transaction, error)
	FindByMedicalRecordNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) ([]Transaction, error)
	Insert(ctx context.Context, tx *sql.Tx, transaction Transaction) (Transaction, error)
	Update(ctx context.Context, tx *sql.Tx, transaction Transaction) error
}

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (t *TransactionRepositoryImpl) FindByMedicalRecordNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) ([]Transaction, error) {
	query := `SELECT TR.transaction_id, TR.register_id, TR.registration_fee, TR.examination_fee, TR.total_fee, TR.discount, TR.total_payment, TR.payment_type, TR.payment_status, TR.created_at, TR.updated_at
			FROM transactions as TR
			INNER JOIN register AS RG ON TR.register_id = RG.register_id
			INNER JOIN patient AS PT ON RG.medical_record_no = PT.medical_record_no
			WHERE PT.medical_record_no = ?`

	rows, err := tx.QueryContext(ctx, query, medicalRecordNo)
	if err != nil {
		return []Transaction{}, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.TransactionID, &transaction.RegisterID, &transaction.RegistrationFee, &transaction.ExaminationFee, &transaction.TotalFee, &transaction.Discount, &transaction.TotalPayment, &transaction.PaymentType, &transaction.PaymentStatus, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return []Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Transaction, error) {
	query := "SELECT transaction_id, register_id, registration_fee, examination_fee, total_fee, discount, total_payment, payment_type, payment_status, created_at, updated_at FROM transactions"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Transaction{}, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.TransactionID, &transaction.RegisterID, &transaction.RegistrationFee, &transaction.ExaminationFee, &transaction.TotalFee, &transaction.Discount, &transaction.TotalPayment, &transaction.PaymentType, &transaction.PaymentStatus, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return []Transaction{}, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionId int) (Transaction, error) {
	query := "SELECT transaction_id, register_id, registration_fee, examination_fee, total_fee, discount, total_payment, payment_type, payment_status, created_at, updated_at FROM transactions WHERE transaction_id = ?"
	row, err := tx.QueryContext(ctx, query, transactionId)
	if err != nil {
		return Transaction{}, err
	}
	defer row.Close()

	var transaction Transaction
	if row.Next() {
		if err := row.Scan(&transaction.TransactionID, &transaction.RegisterID, &transaction.RegistrationFee, &transaction.ExaminationFee, &transaction.TotalFee, &transaction.Discount, &transaction.TotalPayment, &transaction.PaymentType, &transaction.PaymentStatus, &transaction.CreatedAt, &transaction.UpdatedAt); err != nil {
			return Transaction{}, err
		}

		return transaction, nil
	}

	return Transaction{}, custom.ErrNotFound
}

func (t *TransactionRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, transaction Transaction) (Transaction, error) {
	query := "INSERT INTO transactions(register_id, registration_fee, examination_fee, total_fee, discount, total_payment, payment_type, payment_status) VALUES (?,?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, query, transaction.RegisterID, transaction.RegistrationFee, transaction.ExaminationFee, transaction.TotalFee, transaction.Discount, transaction.TotalPayment, transaction.PaymentType, transaction.PaymentStatus)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (t *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction Transaction) error {
	query := "UPDATE transactions SET registration_fee=?,examination_fee=?,total_fee=?,discount=?,total_payment=?,payment_type=?,payment_status=? WHERE transaction_id = ?"
	_, err := tx.ExecContext(ctx, query, transaction.RegistrationFee, transaction.ExaminationFee, transaction.TotalFee, transaction.Discount, transaction.TotalPayment, transaction.PaymentType, transaction.PaymentStatus, transaction.TransactionID)
	if err != nil {
		return err
	}

	return nil
}
