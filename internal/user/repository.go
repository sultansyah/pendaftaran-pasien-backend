package user

import (
	"context"
	"database/sql"
	"pendaftaran-pasien-backend/internal/custom"
)

type UserRepository interface {
	FindByCodeNamePassword(ctx context.Context, tx *sql.Tx, user User) (User, error)
	Update(ctx context.Context, tx *sql.Tx, user User) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) FindByCodeNamePassword(ctx context.Context, tx *sql.Tx, user User) (User, error) {
	sql := "select id, staff_code, staff_name, password, created_at, updated_at from users where staff_code = ? && staff_name = ? && password = ?"
	row, err := tx.QueryContext(ctx, sql, user.StaffCode, user.StaffName, user.Password)
	if err != nil {
		return User{}, err
	}
	defer row.Close()

	if row.Next() {
		err := row.Scan(&user.Id, &user.StaffCode, &user.StaffName, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return User{}, err
		}

		return user, nil
	}

	return User{}, custom.ErrNotFound
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user User) error {
	sql := "update users set staff_code = ?, staff_name = ?, password = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, user.StaffCode, user.StaffName, user.Password, user.Id)
	if err != nil {
		return err
	}

	return nil
}
