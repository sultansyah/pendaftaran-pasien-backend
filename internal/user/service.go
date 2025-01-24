package user

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
	"pendaftaran-pasien-backend/internal/helper"
	loginhistory "pendaftaran-pasien-backend/internal/login_history"
	"pendaftaran-pasien-backend/internal/token"
)

type UserService interface {
	Login(ctx context.Context, input LoginUserInput) (User, string, error)
	UpdatePassword(ctx context.Context, input UpdatePasswordUserInput) error
}

type UserServiceImpl struct {
	DB                     *sql.DB
	UserRepository         UserRepository
	TokenService           token.TokenService
	LoginHistoryRepository loginhistory.LoginHistoryRepository
}

func NewUserService(DB *sql.DB, userRepository UserRepository, tokenService token.TokenService, loginHistoryRepository loginhistory.LoginHistoryRepository) UserService {
	return &UserServiceImpl{
		DB:                     DB,
		UserRepository:         userRepository,
		TokenService:           tokenService,
		LoginHistoryRepository: loginHistoryRepository,
	}
}

func (u *UserServiceImpl) Login(ctx context.Context, input LoginUserInput) (User, string, error) {
	isLoginSuccess := false

	tx, err := u.DB.Begin()
	if err != nil {
		return User{}, "", err
	}
	defer helper.HandleTransaction(tx, &err)

	user := User{
		StaffName: input.StaffName,
		StaffCode: input.StaffCode,
		Password:  input.Password,
	}

	user, err = u.UserRepository.FindByCodeNamePassword(ctx, tx, user)
	if err != nil {
		return User{}, "", err
	}
	if user.Id <= 0 {
		return User{}, "", custom.ErrNotFound
	}

	isSame, err := helper.ComparePassword(user.Password, input.Password)
	if err != nil {
		return User{}, "", err
	}
	if !isSame {
		return User{}, "", custom.ErrUnauthorized
	}

	token, err := u.TokenService.GenerateToken(user.Id)
	if err != nil {
		return User{}, "", err
	}

	defer func() {
		loginHistory := loginhistory.LoginHistory{
			UserId:    user.Id,
			LoginTime: input.Date,
			Success:   isLoginSuccess,
		}

		_, err = u.LoginHistoryRepository.Insert(ctx, tx, loginHistory)
	}()

	isLoginSuccess = true

	return user, token, nil
}

func (u *UserServiceImpl) UpdatePassword(ctx context.Context, input UpdatePasswordUserInput) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.HandleTransaction(tx, &err)

	user := User{
		StaffName: input.StaffName,
		StaffCode: input.StaffCode,
		Password:  input.Password,
	}

	user, err = u.UserRepository.FindByCodeNamePassword(ctx, tx, user)
	if err != nil {
		return err
	}
	if user.Id <= 0 {
		return custom.ErrNotFound
	}

	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	err = u.UserRepository.Update(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}
