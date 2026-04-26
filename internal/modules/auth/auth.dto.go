package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Role string `json:"role,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserInfoResponse struct {
	ID    uint `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}