package patient

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type PatientHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PatientHandlerImpl struct {
	PatientService PatientService
}

func NewPatientHandler(patientService PatientService) PatientHandler {
	return &PatientHandlerImpl{PatientService: patientService}
}

func (p *PatientHandlerImpl) Create(c *gin.Context) {
	var input CreatePatientInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	patient, err := p.PatientService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success create data patient",
		Data:    patient,
	})
}

func (p *PatientHandlerImpl) Delete(c *gin.Context) {
	var input GetPatientInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	err := p.PatientService.Delete(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success delete data patient",
		Data:    "OK",
	})
}

func (p *PatientHandlerImpl) GetAll(c *gin.Context) {
	var input GetPatientInput
	if !helper.BindAndValidate(c, &input, "form") {
		return
	}

	patients, err := p.PatientService.GetAll(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data patient",
		Data:    patients,
	})
}

func (p *PatientHandlerImpl) Update(c *gin.Context) {
	var inputId GetPatientInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData CreatePatientInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := p.PatientService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update data patient",
		Data:    "OK",
	})
}
