package user

type LoginUserInput struct {
	StaffCode string `json:"staff_code" binding:"required"`
	StaffName string `json:"staff_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Date      string `json:"date" binding:"required"`
}

type UpdatePasswordUserInput struct {
	StaffCode string `json:"staff_code" binding:"required"`
	StaffName string `json:"staff_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
