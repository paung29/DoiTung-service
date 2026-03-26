package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	jwtService "github.com/doitung/DoiTung-service/internal/common/jwt"
)


func (s *service) Login(form LoginRequest) (string, AuthResponse, error) {

	account, err := s.accountRepo.FindByEmail(form.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", AuthResponse{
				Success: false,
				Message: "account not found",
			}, nil
		}

		return "", AuthResponse{
			Success: false,
			Message: "internal server error",
		}, err
	}
		
	err = bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(form.Password),
	)

	if err != nil {
		return "", AuthResponse{
			Success: false,
			Message: "invalid password",
		}, nil
	}

	token, err := jwtService.GenerateToken(account.AccountID, string(account.Role))

	if err != nil {
		return "", AuthResponse{
			Success: false,
			Message: "failed to generate token",
		}, err
	}

	return token, AuthResponse{
		Success: true,
		Message: "login successful",
	}, nil

}