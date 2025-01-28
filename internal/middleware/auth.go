package middleware

import (
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
	"pendaftaran-pasien-backend/internal/token"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(tokenService token.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil access token dari cookie
		tokenString, err := c.Cookie("auth_token")
		if err != nil {
			// jika tidak ada access token, kirimkan error
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		// validasi access token
		token, err := tokenService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		// jika access token valid, cek expiration
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		// cek apakah access token sudah expired
		if time.Now().Unix() > int64(exp) {
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		// ambil user_id dari access token
		userId, ok := claims["user_id"].(float64)
		if !ok {
			c.SetCookie("auth_token", "", -1, "/", "", false, true)
			helper.HandleErrorResponde(c, custom.ErrUnauthorized)
			c.Abort()
			return
		}

		// set userId pada context
		c.Set("userId", int(userId))

		c.Next()
	}
}
