package queue

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
	"pendaftaran-pasien-backend/internal/patient"
)

type QueueService interface {
	GetAll(ctx context.Context, input GetQueueInput) ([]Queue, error)
	GetById(ctx context.Context, input GetQueueByIdInput) (Queue, error)
	Update(ctx context.Context, inputId GetQueueByIdInput, inputData UpdateQueueInput) error
}

type QueueServiceImpl struct {
	DB                *sql.DB
	QueueRepository   QueueRepository
	PatientRepository patient.PatientRepository
}

func NewQueueService(DB *sql.DB, queueRepository QueueRepository, patientRepository patient.PatientRepository) QueueService {
	return &QueueServiceImpl{
		DB:                DB,
		QueueRepository:   queueRepository,
		PatientRepository: patientRepository,
	}
}

func (t *QueueServiceImpl) GetAll(ctx context.Context, input GetQueueInput) ([]Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	filters := helper.GenerateFilter(input)

	if medicalRecordNo, ok := filters["medical_record_no"]; ok {
		medicalRecordNo, ok := medicalRecordNo.(string)
		if !ok {
			return []Queue{}, custom.ErrInternal
		}

		_, err := t.PatientRepository.FindByNoMR(ctx, tx, medicalRecordNo)
		if err != nil {
			return []Queue{}, err
		}
	}

	queues, err := t.QueueRepository.FindQueues(ctx, tx, filters)
	if err != nil {
		return []Queue{}, err
	}

	return queues, nil
}

func (t *QueueServiceImpl) GetById(ctx context.Context, input GetQueueByIdInput) (Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	queue, err := t.QueueRepository.FindById(ctx, tx, input.QueueID)
	if err != nil {
		return Queue{}, err
	}
	if queue.QueueID <= 0 {
		return Queue{}, custom.ErrNotFound
	}

	return queue, nil
}

func (t *QueueServiceImpl) Update(ctx context.Context, inputId GetQueueByIdInput, inputData UpdateQueueInput) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	queue, err := t.QueueRepository.FindById(ctx, tx, inputId.QueueID)
	if err != nil {
		return err
	}
	if queue.QueueID <= 0 {
		return custom.ErrNotFound
	}

	queue.RegisterID = inputData.RegisterID
	queue.QueueNumber = inputData.QueueNumber

	err = t.QueueRepository.Update(ctx, tx, queue)
	if err != nil {
		return err
	}

	return nil
}
