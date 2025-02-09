package custom

import (
	"errors"
)

var (
	ErrAlreadyExists               = errors.New("resource already exists")
	ErrIdentityNumberAlreadyExists = errors.New("identity number already used")
	ErrNotFound                    = errors.New("resource not found")
	ErrIdentityNumberNotFound      = errors.New("identity number not found")
	ErrRegisterNotFound            = errors.New("register number not found")
	ErrPolyclinicNotFound          = errors.New("polyclinic not found")
	ErrDoctorNotFound              = errors.New("doctor not found")
	ErrMedicalRecordNotFound       = errors.New("medical record not found")
	ErrInternal                    = errors.New("internal server error")
	ErrUnauthorized                = errors.New("unauthorized")
	ErrForbidden                   = errors.New("you are not authorized to access this resource")
	ErrImageRequired               = errors.New("image is required")
	ErrConflict                    = errors.New("duplicate entry for key")
	ErrInsufficientStock           = errors.New("insufficient stock")
	ErrInvalidCredentials          = errors.New("name, code, or password is incorrect")
)
