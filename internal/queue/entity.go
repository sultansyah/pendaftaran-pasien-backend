package queue

import "time"

type Queue struct {
	QueueID     int       `json:"queue_id"`
	RegisterID  string    `json:"register_id"`
	QueueNumber int       `json:"queue_number"`
	IsCompleted int       `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
