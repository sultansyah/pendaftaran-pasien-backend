package transaction

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetByMedicalRecordNo(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type TransactionHandlerImpl struct {
	TransactionService TransactionService
}

func NewTransactionHandler(transactionService TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{TransactionService: transactionService}
}

func (t *TransactionHandlerImpl) Create(c *gin.Context) {
	var input CreateTransactionInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	transaction, err := t.TransactionService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success create transaction",
		Data:    transaction,
	})
}

func (t *TransactionHandlerImpl) GetAll(c *gin.Context) {
	transactions, err := t.TransactionService.GetAll(c.Request.Context())
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all transaction",
		Data:    transactions,
	})
}

func (t *TransactionHandlerImpl) GetById(c *gin.Context) {
	var input GetTransactionInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	transaction, err := t.TransactionService.GetById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get transaction",
		Data:    transaction,
	})
}

func (t *TransactionHandlerImpl) GetByMedicalRecordNo(c *gin.Context) {
	var input GetTransactionByMedicalRecordNoInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	transaction, err := t.TransactionService.GetByMedicalRecordNo(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get transaction",
		Data:    transaction,
	})
}

func (t *TransactionHandlerImpl) Update(c *gin.Context) {
	var inputId GetTransactionInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData UpdateTransactionInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := t.TransactionService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update transaction",
		Data:    "OK",
	})
}
