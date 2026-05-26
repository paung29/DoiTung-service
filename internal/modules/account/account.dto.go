package account

type AccountCreateForm struct {
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password" validate:"required,min=6"`
	Role         string  `json:"role" validate:"required,min=5"`
	Name         string  `json:"name" validate:"required,min=2,max=100"`
	PhoneNo      *string `json:"phone_no" validate:"omitempty"`
	ActiveStatus bool    `json:"active_status"`
}

type AccountCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AccountUpdateInfoForm struct {
	UserId       uint    `json:"user_id" validate:"required"`
	Name         *string `json:"name" validate:"omitempty,min=2,max=100"`
	Role         *string `json:"role" validate:"omitempty,min=5"`
	PhoneNo      *string `json:"phone_no" validate:"omitempty,min=8,max=20"`
	ActiveStatus *bool   `json:"active_status"`
}

type AccountUpdateInfoResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AccountPasswordUpdateForm struct {
	UserId   uint   `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type AccountPasswordUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AccountDetails struct {
	UserId       uint    `json:"user_id"`
	Email        string  `json:"email"`
	Name         *string `json:"name"`
	Role         *string `json:"role"`
	PhoneNo      *string `json:"phone_no"`
	ActiveStatus *bool   `json:"active_status"`
}

type AccountLists struct {
	Accounts []AccountDetails `json:"accounts"`
}
