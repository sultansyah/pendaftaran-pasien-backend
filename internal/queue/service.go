package queue

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/helper"
)

type QueueService interface {
	GetAll(ctx context.Context) ([]Queue, error)
	GetAllByDay(ctx context.Context, input GetQueueByDayInput) ([]Queue, error)
	GetAllByMedicalRecordNo(ctx context.Context, input GetQueueByMedicalRecordNoInput) ([]Queue, error)
	GetById(ctx context.Context, input GetQueueInput) (Queue, error)
	Update(ctx context.Context, inputId GetQueueInput, inputData UpdateQueueInput) error
}

type QueueServiceImpl struct {
	DB              *sql.DB
	QueueRepository QueueRepository
}

func NewQueueService(DB *sql.DB, queueRepository QueueRepository) QueueService {
	return &QueueServiceImpl{
		DB:              DB,
		QueueRepository: queueRepository,
	}
}

func (t *QueueServiceImpl) GetAll(ctx context.Context) ([]Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	queues, err := t.QueueRepository.FindAll(ctx, tx)
	if err != nil {
		return []Queue{}, err
	}

	return queues, nil
}

func (t *QueueServiceImpl) GetAllByDay(ctx context.Context, input GetQueueByDayInput) ([]Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	day, err := helper.ConvertDayToEnglish(input.Day)
	if err != nil {
		return []Queue{}, err
	}

	queues, err := t.QueueRepository.FindByDay(ctx, tx, day)
	if err != nil {
		return []Queue{}, err
	}

	return queues, nil
}

func (t *QueueServiceImpl) GetAllByMedicalRecordNo(ctx context.Context, input GetQueueByMedicalRecordNoInput) ([]Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return []Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	queues, err := t.QueueRepository.FindByMedicalRecordNo(ctx, tx, input.MedicalRecordNo)
	if err != nil {
		return []Queue{}, err
	}

	return queues, nil
}

func (t *QueueServiceImpl) GetById(ctx context.Context, input GetQueueInput) (Queue, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return Queue{}, err
	}
	defer helper.HandleTransaction(tx, &err)

	queue, err := t.QueueRepository.FindById(ctx, tx, input.QueueID)
	if err != nil {
		return Queue{}, err
	}

	return queue, nil
}

func (t *QueueServiceImpl) Update(ctx context.Context, inputId GetQueueInput, inputData UpdateQueueInput) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	queue, err := t.QueueRepository.FindById(ctx, tx, inputId.QueueID)
	if err != nil {
		return err
	}

	queue.RegisterID = inputData.RegisterID
	queue.QueueNumber = inputData.QueueNumber

	err = t.QueueRepository.Update(ctx, tx, queue)
	if err != nil {
		return err
	}

	return nil
}
