package account

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
)

func (s *service) CreateAccount(form AccountCreateForm) (AccountCreateResponse, error) {

	role := enums.Role(form.Role)

	switch role {
	case enums.RoleAdmin, enums.RoleStaff:
	default:
		return AccountCreateResponse{
			Success: false,
			Message: "invalid role",
		}, nil
	}

	existingAccount, err := s.accountRepo.FindByEmail(form.Email)

	if err == nil && existingAccount != nil {
		return AccountCreateResponse{
			Success: false,
			Message: "email already exists",
		}, nil
	}

	hashedPassword, err := utils.HashedPassword(form.Password)

	if err != nil {
		return  AccountCreateResponse{
			Success: false,
			Message: "failed to hash password",
		}, err
	}


	account := &models.Account{
		Email: form.Email,
		PasswordHash: hashedPassword,
		Role: role,
	}

	if err := s.accountRepo.Create(account); err != nil {
		return AccountCreateResponse{
			Success: false,
			Message: "failed to create account",
		}, err
	}

	return AccountCreateResponse{
		Success: true,
		Message: "account created successfully",
	}, nil
}