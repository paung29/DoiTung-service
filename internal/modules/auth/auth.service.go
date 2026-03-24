package auth

import "github.com/doitung/DoiTung-service/internal/modules/account"



type AuthService interface {
	 Login(form LoginRequest) (string,  AuthResponse, error)
}

type service struct {
	accountRepo account.AccountRepository 
}

func NewAuthService(accountRepo account.AccountRepository) AuthService {
	return &service{
		accountRepo: accountRepo,
	}
}



