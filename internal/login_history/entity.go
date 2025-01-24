package loginhistory

import "time"

type LoginHistory struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	LoginTime time.Time `json:"login_time"`
	Success   bool      `json:"success"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
