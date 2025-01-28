package token

import (
	"pendaftaran-pasien-backend/internal/custom"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateToken(userId int) (string, string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type TokenServiceImpl struct {
	Key []byte
}

func NewTokenService(key []byte) TokenService {
	return &TokenServiceImpl{Key: key}
}

func (t *TokenServiceImpl) GenerateToken(userId int) (string, string, error) {
	// berlaku 3 hari
	accessClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Second * 60).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(t.Key)
	if err != nil {
		return "", "", err
	}

	// berlaku 30 hari
	refreshClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(t.Key)
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

func (t *TokenServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, custom.ErrUnauthorized
		}
		return t.Key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, custom.ErrUnauthorized
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, custom.ErrUnauthorized
	}

	if time.Now().Unix() > int64(exp) {
		return nil, custom.ErrUnauthorized
	}

	return token, nil
}
