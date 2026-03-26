package account

type AccountCreateForm struct {
	Email string	`json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role string `json:"role" validate:"required,min=5"`
}

type AccountCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}