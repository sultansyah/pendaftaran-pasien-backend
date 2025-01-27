package queue

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type QueueRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]Queue, error)
	FindById(ctx context.Context, tx *sql.Tx, queueId int) (Queue, error)
	FindByDay(ctx context.Context, tx *sql.Tx, day string) ([]Queue, error)
	FindByMedicalRecordNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) ([]Queue, error)
	Insert(ctx context.Context, tx *sql.Tx, queue Queue) (Queue, error)
	Update(ctx context.Context, tx *sql.Tx, queue Queue) error
	CountQueueToday(ctx context.Context, tx *sql.Tx, day string) (int, error)
}

type QueueRepositoryImpl struct {
}

func NewQueueRepository() QueueRepository {
	return &QueueRepositoryImpl{}
}

func (t *QueueRepositoryImpl) CountQueueToday(ctx context.Context, tx *sql.Tx, day string) (int, error) {
	query := "SELECT COUNT(queue_id) FROM queue WHERE DAYNAME(created_at) = ?"
	row, err := tx.QueryContext(ctx, query, day)
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

func (t *QueueRepositoryImpl) FindByDay(ctx context.Context, tx *sql.Tx, day string) ([]Queue, error) {
	query := "SELECT queue_id, register_id, queue_number, created_at, updated_at FROM queue WHERE DAYNAME(created_at) = ?"

	rows, err := tx.QueryContext(ctx, query, day)
	if err != nil {
		return []Queue{}, err
	}
	defer rows.Close()

	var queues []Queue
	for rows.Next() {
		var queue Queue
		if err := rows.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return []Queue{}, err
		}

		queues = append(queues, queue)
	}

	return queues, nil
}

func (t *QueueRepositoryImpl) FindByMedicalRecordNo(ctx context.Context, tx *sql.Tx, medicalRecordNo string) ([]Queue, error) {
	query := `SELECT QU.queue_id, QU.register_id, QU.queue_number, QU.created_at, QU.updated_at
FROM queue as QU
INNER JOIN register AS RG ON QU.register_id = RG.register_id
INNER JOIN patient AS PT ON RG.medical_record_no = PT.medical_record_no
WHERE PT.medical_record_no = ?`

	rows, err := tx.QueryContext(ctx, query, medicalRecordNo)
	if err != nil {
		return []Queue{}, err
	}
	defer rows.Close()

	var queues []Queue
	for rows.Next() {
		var queue Queue
		if err := rows.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return []Queue{}, err
		}

		queues = append(queues, queue)
	}

	return queues, nil
}

func (t *QueueRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]Queue, error) {
	query := "SELECT queue_id, register_id, queue_number, created_at, updated_at FROM queue"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []Queue{}, err
	}
	defer rows.Close()

	var queues []Queue
	for rows.Next() {
		var queue Queue
		if err := rows.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return []Queue{}, err
		}

		queues = append(queues, queue)
	}

	return queues, nil
}

func (t *QueueRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, queueId int) (Queue, error) {
	query := "SELECT queue_id, register_id, queue_number, created_at, updated_at FROM queue WHERE queue_id = ?"
	row, err := tx.QueryContext(ctx, query, queueId)
	if err != nil {
		return Queue{}, err
	}
	defer row.Close()

	var queue Queue
	if row.Next() {
		if err := row.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return Queue{}, err
		}

		return queue, nil
	}

	return Queue{}, nil
}

func (t *QueueRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, queue Queue) (Queue, error) {
	query := "INSERT INTO queue(register_id, queue_number) VALUES (?,?)"
	_, err := tx.ExecContext(ctx, query, queue.RegisterID, queue.QueueNumber)
	if err != nil {
		return Queue{}, err
	}

	return queue, nil
}

func (t *QueueRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, queue Queue) error {
	query := "UPDATE queue SET register_id=?,queue_number=? WHERE queue_id = ?"
	_, err := tx.ExecContext(ctx, query, queue.RegisterID, queue.QueueNumber)
	if err != nil {
		return err
	}

	return nil
}
