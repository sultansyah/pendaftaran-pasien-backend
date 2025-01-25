package polyclinic

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type PolyclinicHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PolyclinicHandlerImpl struct {
	PolyclinicService PolyclinicService
}

func NewPolyclinicHandler(polyclinicService PolyclinicService) PolyclinicHandler {
	return &PolyclinicHandlerImpl{PolyclinicService: polyclinicService}
}

func (p *PolyclinicHandlerImpl) Create(c *gin.Context) {
	var input CreatePolyclinicInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	polyclinic, err := p.PolyclinicService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success create polyclinic",
		Data:    polyclinic,
	})
}

func (p *PolyclinicHandlerImpl) Delete(c *gin.Context) {
	var input GetPolyclinicInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	err := p.PolyclinicService.Delete(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success delete polyclinic",
		Data:    "OK",
	})
}

func (p *PolyclinicHandlerImpl) GetAll(c *gin.Context) {
	polyclinics, err := p.PolyclinicService.GetAll(c.Request.Context())
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get all data polyclinic",
		Data:    polyclinics,
	})
}

func (p *PolyclinicHandlerImpl) GetById(c *gin.Context) {
	var input GetPolyclinicInput
	if !helper.BindAndValidate(c, &input, "uri") {
		return
	}

	polyclinic, err := p.PolyclinicService.GetById(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success get data polyclinic",
		Data:    polyclinic,
	})
}

func (p *PolyclinicHandlerImpl) Update(c *gin.Context) {
	var inputId GetPolyclinicInput
	if !helper.BindAndValidate(c, &inputId, "uri") {
		return
	}

	var inputData CreatePolyclinicInput
	if !helper.BindAndValidate(c, &inputData, "json") {
		return
	}

	err := p.PolyclinicService.Update(c.Request.Context(), inputId, inputData)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update data polyclinic",
		Data:    "OK",
	})
}
