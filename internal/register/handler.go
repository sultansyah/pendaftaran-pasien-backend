package register

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type RegisterHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetLatestByMRNo(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type RegisterHandlerImpl struct {
	RegisterService RegisterService
}

func NewRegisterHandler(registerService RegisterService) RegisterHandler {
	return &RegisterHandlerImpl{RegisterService: registerService}
}

func (r *RegisterHandlerImpl) Create(c *gin.Context) {
	var input CreateRegisterInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	register, err := r.RegisterService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success register patient",
		Data:    register,
	})
}

func (r *RegisterHandlerImpl) Delete(c *gin.Context) {
	var input GetRegisterInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	err := r.RegisterService.Delete(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success delete data register",
		Data:    "OK",
	})
}

func (r *RegisterHandlerImpl) GetAll(c *gin.Context) {
	registers, err := r.RegisterService.GetAll(c.Request.Context())
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data register",
		Data:    registers,
	})
}

func (r *RegisterHandlerImpl) GetById(c *gin.Context) {
	var input GetRegisterInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	register, err := r.RegisterService.GetById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data register",
		Data:    register,
	})
}

func (r *RegisterHandlerImpl) GetLatestByMRNo(c *gin.Context) {
	var input GetRegisterByMRNoInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	register, err := r.RegisterService.GetLatestByMRNo(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get latest data register",
		Data:    register,
	})
}

func (r *RegisterHandlerImpl) Update(c *gin.Context) {
	var inputId GetRegisterInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData CreateRegisterInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := r.RegisterService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update data register",
		Data:    "OK",
	})
}
