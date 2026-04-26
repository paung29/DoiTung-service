package auth

import (
	"errors"

	jwtService "github.com/doitung/DoiTung-service/internal/common/jwt"
	"github.com/doitung/DoiTung-service/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func (s *service) Login(form LoginRequest) (string, AuthResponse, error) {

	account, err := s.accountRepo.FindByEmail(form.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", AuthResponse{}, utils.UnauthorizedError("account not found")
		}

		return "", AuthResponse{}, utils.SystemError("failed to query account")
	}
		
	err = bcrypt.CompareHashAndPassword(
		[]byte(account.PasswordHash),
		[]byte(form.Password),
	)

	if err != nil {
		return "", AuthResponse{}, utils.UnauthorizedError("invalid password")
	}

	token, err := jwtService.GenerateToken(account.AccountID, string(account.Role))

	if err != nil {
		return "", AuthResponse{}, utils.SystemError("failed to generate token")
	}

	return token, AuthResponse{
		Role: string(account.Role),
		Success: true,
		Message: "login successful",
	}, nil

}

func (s *service) GetUserInfo(userId uint) (UserInfoResponse, error) {

	account, err := s.accountRepo.FindByID(userId)
	if err != nil {
		return UserInfoResponse{}, utils.SystemError("failed to query user info")
	}

	return UserInfoResponse{
		ID:    account.AccountID,
		Email: account.Email,
		Role:  string(account.Role),
	}, nil
}
