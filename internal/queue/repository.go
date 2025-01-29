package queue

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"strings"
)

type QueueRepository interface {
	FindQueues(ctx context.Context, tx *sql.Tx, filters map[string]any) ([]Queue, error)
	FindById(ctx context.Context, tx *sql.Tx, queueId int) (Queue, error)
	Insert(ctx context.Context, tx *sql.Tx, queue Queue) (Queue, error)
	Update(ctx context.Context, tx *sql.Tx, queue Queue) error
	CountQueueByDay(ctx context.Context, tx *sql.Tx, date string) (int, error)
	SetCompleted(ctx context.Context, tx *sql.Tx, id int) error
}

type QueueRepositoryImpl struct {
}

func NewQueueRepository() QueueRepository {
	return &QueueRepositoryImpl{}
}

func (t *QueueRepositoryImpl) CountQueueByDay(ctx context.Context, tx *sql.Tx, date string) (int, error) {
	query := "SELECT COALESCE(MAX(queue_number), 1) FROM queue WHERE DATE(created_at) = ?"
	row, err := tx.QueryContext(ctx, query, date)
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

func (t *QueueRepositoryImpl) FindQueues(ctx context.Context, tx *sql.Tx, filters map[string]any) ([]Queue, error) {
	query := "SELECT QU.queue_id, QU.register_id, QU.queue_number, QU.is_completed, QU.created_at, QU.updated_at FROM queue AS QU"
	join := ""
	whereConditions := []string{}
	args := []any{}

	// filters
	if medicalRecordNo, ok := filters["medical_record_no"]; ok {
		join += " INNER JOIN register AS RG ON QU.register_id = RG.register_id" +
			" INNER JOIN patient AS PT ON RG.medical_record_no = PT.medical_record_no"
		whereConditions = append(whereConditions, "PT.medical_record_no = ?")
		args = append(args, medicalRecordNo)
	}

	if date, ok := filters["date"]; ok {
		whereConditions = append(whereConditions, "DATE(QU.created_at) = ?")
		args = append(args, date)
	}

	// combine where conditions
	where := ""
	if len(whereConditions) > 0 {
		where = " WHERE " + strings.Join(whereConditions, " AND ")
	}

	// combine all
	query = query + join + where

	// execute query
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return []Queue{}, err
	}
	defer rows.Close()

	var queues []Queue
	for rows.Next() {
		var queue Queue
		if err := rows.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.IsCompleted, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return []Queue{}, err
		}
		queues = append(queues, queue)
	}

	return queues, nil
}

func (t *QueueRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, queueId int) (Queue, error) {
	query := "SELECT queue_id, register_id, queue_number, is_completed, created_at, updated_at FROM queue WHERE queue_id = ?"
	row, err := tx.QueryContext(ctx, query, queueId)
	if err != nil {
		return Queue{}, err
	}
	defer row.Close()

	var queue Queue
	if row.Next() {
		if err := row.Scan(&queue.QueueID, &queue.RegisterID, &queue.QueueNumber, &queue.IsCompleted, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
			return Queue{}, err
		}

		return queue, nil
	}

	return Queue{}, nil
}

func (t *QueueRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, queue Queue) (Queue, error) {
	query := "INSERT INTO queue(register_id, queue_number) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, query, queue.RegisterID, queue.QueueNumber)
	if err != nil {
		return Queue{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Queue{}, err
	}

	queue.QueueID = int(id)

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

func (t *QueueRepositoryImpl) SetCompleted(ctx context.Context, tx *sql.Tx, id int) error {
	query := "UPDATE queue SET is_completed=1 WHERE queue_id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
