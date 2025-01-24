package loginhistory

import (
	"context"
	"database/sql"
)

type LoginHistoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, loginHistory LoginHistory) (LoginHistory, error)
}

type LoginHistoryRepositoryImpl struct {
}

func NewLoginHistoryRepository() LoginHistoryRepository {
	return &LoginHistoryRepositoryImpl{}
}

func (r *LoginHistoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, loginHistory LoginHistory) (LoginHistory, error) {
	sql := "insert into login_history(user_id, login_time, success) VALUES(?,?,?)"
	result, err := tx.ExecContext(ctx, sql, loginHistory.UserId, loginHistory.LoginTime, loginHistory.Success)
	if err != nil {
		return loginHistory, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return loginHistory, err
	}

	loginHistory.Id = int(id)
	return loginHistory, nil
}
