package account

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

func (s *service) CreateAccount(form AccountCreateForm) (AccountCreateResponse, error) {

	role := enums.Role(form.Role)

	switch role {
	case enums.RoleAdmin, enums.RoleStaff:
	default:
		return AccountCreateResponse{}, utils.BadRequestError("invalid role")
	}

	existingAccount, err := s.accountRepo.FindByEmail(form.Email)

	if err == nil && existingAccount != nil {
		return AccountCreateResponse{}, utils.BadRequestError("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return AccountCreateResponse{}, utils.SystemError("failed to check existing account")
	}

	hashedPassword, err := utils.HashedPassword(form.Password)

	if err != nil {
		return AccountCreateResponse{}, utils.SystemError("failed to hash password")
	}

	account := &models.Account{
		Email:        form.Email,
		PasswordHash: hashedPassword,
		Role:         role,
	}

	if err := s.accountRepo.Create(account); err != nil {
		return AccountCreateResponse{}, utils.SystemError("failed to create account")
	}

	return AccountCreateResponse{
		Success: true,
		Message: "account created successfully",
	}, nil
}

func (s *service) UpdateAccount(form AccountUpdateForm) (AccountUpdateResponse, error) {
	account, err := s.accountRepo.FindByID(form.UserId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return AccountUpdateResponse{}, utils.BadRequestError("account not found")
		}
		return AccountUpdateResponse{}, utils.SystemError("failed to find account")
	}

	if form.Name != nil {
		account.Name = *form.Name
	}

	if form.PhoneNo != nil {
		account.PhoneNo = *form.PhoneNo
	}

	if form.Password != nil {
		hashedPassword, err := utils.HashedPassword(*form.Password)
		if err != nil {
			return AccountUpdateResponse{}, utils.SystemError("failed to hash password")
		}
		account.PasswordHash = hashedPassword
	}

	if form.ActiveStatus != nil {
		account.ActiveStatus = *form.ActiveStatus
	}

	if err := s.accountRepo.Update(account); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return AccountUpdateResponse{}, utils.BadRequestError("account not found")
		}
		return AccountUpdateResponse{}, utils.SystemError("failed to update account")
	}

	return AccountUpdateResponse{
		Success: true,
		Message: "account updated successfully",
	}, nil
}
