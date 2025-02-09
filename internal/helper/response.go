package helper

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/custom"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func APIResponse(c *gin.Context, response WebResponse) {
	c.JSON(response.Code, response)
}

func HandleErrorResponde(c *gin.Context, err error) {
	webResponse := WebResponse{
		Data: nil,
	}

	switch err {
	case custom.ErrAlreadyExists:
		webResponse.Code = http.StatusConflict
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrPolyclinicNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrMedicalRecordNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrDoctorNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrRegisterNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrIdentityNumberNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrIdentityNumberAlreadyExists:
		webResponse.Code = http.StatusConflict
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrInternal:
		webResponse.Code = http.StatusInternalServerError
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case bcrypt.ErrMismatchedHashAndPassword:
		webResponse.Code = http.StatusUnauthorized
		webResponse.Status = "unauthorized"
		webResponse.Message = custom.ErrInvalidCredentials.Error()
	case custom.ErrUnauthorized:
		webResponse.Code = http.StatusUnauthorized
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrImageRequired:
		webResponse.Code = http.StatusBadRequest
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrInsufficientStock:
		webResponse.Code = http.StatusConflict
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case custom.ErrInvalidCredentials:
		webResponse.Code = http.StatusUnauthorized
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	default:
		webResponse.Code = http.StatusInternalServerError
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	}

	APIResponse(c, webResponse)
}
