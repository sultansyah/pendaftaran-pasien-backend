package user

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(c *gin.Context)
	UpdatePassword(c *gin.Context)
}

type UserHandlerImpl struct {
	UserService UserService
}

func NewUserHandler(userService UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func (u *UserHandlerImpl) Login(c *gin.Context) {
	var input LoginUserInput

	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	user, token, err := u.UserService.Login(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	c.SetCookie("auth_token", token, 3600, "/", "", false, true)

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "login success",
		Data:    user,
	})
}

func (u *UserHandlerImpl) UpdatePassword(c *gin.Context) {
	var input UpdatePasswordUserInput
	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	err := u.UserService.UpdatePassword(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success update password",
		Data:    "OK",
	})
}
