package user

import "time"

type User struct {
	Id        int       `json:"id"`
	StaffName string    `json:"staff_name"`
	StaffCode string    `json:"staff_code"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
