package queue

type UpdateQueueInput struct {
	RegisterID  string `json:"register_id" binding:"required"`
	QueueNumber int    `json:"queue_number" binding:"required"`
}

type GetQueueByIdInput struct {
	QueueID int `uri:"queue_id" binding:"required"`
}

type GetQueueInput struct {
	MedicalRecordNo string `form:"medical_record_no"`
	Date            string `form:"date"`
}

type GetQueueByDayInput struct {
}
