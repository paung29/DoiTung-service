package auth

import (
	"github.com/doitung/DoiTung-service/internal/config"
	"github.com/doitung/DoiTung-service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(form LoginRequest) (string, AuthResponse, error) {

	var account models.Account

	err := config.DB.Where("email = ?", form.Email).First(&account).Error

	if err != nil {
		return "", AuthResponse{
			Success: false,
			Message: "account not found",
		}, nil
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

	token, err := GenerateToken(account.AccountID)

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