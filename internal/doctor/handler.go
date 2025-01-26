package doctor

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type DoctorHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetByClinicId(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type DoctorHandlerImpl struct {
	DoctorService DoctorService
}

func NewDoctorHandler(doctorService DoctorService) DoctorHandler {
	return &DoctorHandlerImpl{DoctorService: doctorService}
}

func (p *DoctorHandlerImpl) Create(c *gin.Context) {
	var input CreateDoctorInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	doctor, err := p.DoctorService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success create doctor",
		Data:    doctor,
	})
}

func (p *DoctorHandlerImpl) Delete(c *gin.Context) {
	var input GetDoctorInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	err := p.DoctorService.Delete(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success delete doctor",
		Data:    "OK",
	})
}

func (p *DoctorHandlerImpl) GetAll(c *gin.Context) {
	doctors, err := p.DoctorService.GetAll(c.Request.Context())
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data doctor",
		Data:    doctors,
	})
}

func (p *DoctorHandlerImpl) GetById(c *gin.Context) {
	var input GetDoctorInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	doctor, err := p.DoctorService.GetById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data doctor",
		Data:    doctor,
	})
}

func (p *DoctorHandlerImpl) GetByClinicId(c *gin.Context) {
	var input GetDoctorByClinicInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	doctors, err := p.DoctorService.GetByClinicId(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data doctors",
		Data:    doctors,
	})
}

func (p *DoctorHandlerImpl) Update(c *gin.Context) {
	var inputId GetDoctorInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData CreateDoctorInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := p.DoctorService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update data doctor",
		Data:    "OK",
	})
}
