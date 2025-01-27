package queue

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type QueueHandler interface {
	GetAll(c *gin.Context)
	GetAllByDay(c *gin.Context)
	GetAllByMedicalRecordNo(c *gin.Context)
	GetById(c *gin.Context)
	Update(c *gin.Context)
}

type QueueHandlerImpl struct {
	QueueService QueueService
}

func NewQueueHandler(queueService QueueService) QueueHandler {
	return &QueueHandlerImpl{QueueService: queueService}
}

func (q *QueueHandlerImpl) GetAll(c *gin.Context) {
	queues, err := q.QueueService.GetAll(c.Request.Context())
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data queue",
		Data:    queues,
	})
}

func (q *QueueHandlerImpl) GetAllByDay(c *gin.Context) {
	var input GetQueueByDayInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	queues, err := q.QueueService.GetAllByDay(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data queue",
		Data:    queues,
	})
}

func (q *QueueHandlerImpl) GetAllByMedicalRecordNo(c *gin.Context) {
	var input GetQueueByMedicalRecordNoInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	queues, err := q.QueueService.GetAllByMedicalRecordNo(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data queue",
		Data:    queues,
	})
}

func (q *QueueHandlerImpl) GetById(c *gin.Context) {
	var input GetQueueInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	queue, err := q.QueueService.GetById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data queue",
		Data:    queue,
	})
}

func (q *QueueHandlerImpl) Update(c *gin.Context) {
	var inputId GetQueueInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData UpdateQueueInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := q.QueueService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data queue",
		Data:    "OK",
	})
}
