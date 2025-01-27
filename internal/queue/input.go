package queue

type UpdateQueueInput struct {
	RegisterID  string `json:"register_id" binding:"required"`
	QueueNumber int    `json:"queue_number" binding:"required"`
}

type GetQueueInput struct {
	QueueID int `uri:"queue_id" binding:"required"`
}

type GetQueueByMedicalRecordNoInput struct {
	MedicalRecordNo string `uri:"medical_record_no" binding:"required"`
}

type GetQueueByDayInput struct {
	Day string `uri:"day" binding:"required"`
}
