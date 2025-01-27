package register

import (
	"pendaftaran-pasien-backend/internal/queue"
	"pendaftaran-pasien-backend/internal/transaction"
)

type RegisterWithTransactionAndQueue struct {
	Register    Register                `json:"register"`
	Queue       queue.Queue             `json:"queue"`
	Transaction transaction.Transaction `json:"transaction"`
}

func RegisterWithTransactionAndQueueFormatter(register Register, queue queue.Queue, transaction transaction.Transaction) RegisterWithTransactionAndQueue {
	return RegisterWithTransactionAndQueue{
		Register:    register,
		Queue:       queue,
		Transaction: transaction,
	}
}
