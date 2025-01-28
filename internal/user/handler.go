package user

import (
	"net/http"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
	"pendaftaran-pasien-backend/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHandler interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Logout(c *gin.Context)
	UpdatePassword(c *gin.Context)
}

type UserHandlerImpl struct {
	UserService  UserService
	TokenService token.TokenService
}

func NewUserHandler(userService UserService, tokenService token.TokenService) UserHandler {
	return &UserHandlerImpl{
		UserService:  userService,
		TokenService: tokenService,
	}
}

func (u *UserHandlerImpl) Login(c *gin.Context) {
	var input LoginUserInput

	if !helper.BindAndValidate(c, &input, "json") {
		return
	}

	user, accessToken, refreshToken, err := u.UserService.Login(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	c.SetCookie("auth_token", accessToken, 259200, "/", "", false, true)      // 3 hari
	c.SetCookie("refresh_token", refreshToken, 2592000, "/", "", false, true) // 30 hari

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "login success",
		Data:    user,
	})
}

func (u *UserHandlerImpl) RefreshToken(c *gin.Context) {
	// ambil refresh token dari cookie
	tokenString, err := c.Cookie("refresh_token")
	if err != nil {
		// jika tidak ada refresh token, kirimkan error
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		helper.HandleErrorResponde(c, custom.ErrUnauthorized)
		return
	}

	refreshToken, err := u.TokenService.ValidateToken(tokenString)
	if err != nil {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		helper.HandleErrorResponde(c, err)
		return
	}

	// ambil claims
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		helper.HandleErrorResponde(c, custom.ErrUnauthorized)
		c.Abort()
		return
	}

	// ambil user_id dari refresh token
	userId, ok := claims["user_id"].(float64)
	if !ok {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		helper.HandleErrorResponde(c, custom.ErrUnauthorized)
		c.Abort()
		return
	}

	newAccessToken, _, err := u.TokenService.GenerateToken(int(userId))
	if err != nil {
		helper.HandleErrorResponde(c, err)
		return
	}

	c.SetCookie("auth_token", newAccessToken, 259200, "/", "", false, true) // 3 hari

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "refresh token success",
		Data:    "OK",
	})
}

func (u *UserHandlerImpl) Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success logout",
		Data:    nil,
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
