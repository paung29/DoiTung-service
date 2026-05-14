package account

type AccountCreateForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,min=5"`
}

type AccountCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AccountUpdateForm struct {
	UserId       uint    `json:"user_id" validate:"required"`
	Name         *string `json:"name" validate:"omitempty,min=2,max=100"`
	PhoneNo      *string `json:"phone_no" validate:"omitempty,min=8,max=20"`
	Password     *string `json:"password" validate:"omitempty,min=6"`
	ActiveStatus *bool   `json:"active_status"`
}

type AccountUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
