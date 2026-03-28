package year

import (
	"errors"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

func (s *service) CreateYear(form YearCreateForm) (YearCreateResponse, error){

	existingYear, err := s.yearRepo.FindByYear(form.Year)

	if err == nil && existingYear != nil {
		return YearCreateResponse{}, utils.BadRequestError("year already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return YearCreateResponse{}, utils.SystemError("failed to check existing year")
	}

	tx := s.db.Begin()

	if tx.Error != nil {
		return YearCreateResponse{}, utils.SystemError("failed to start transaction")
	}

	year := &models.Year{
		Year: form.Year,
	}

	if err := s.yearRepo.Create(tx, year); err != nil {
		tx.Rollback()
		return YearCreateResponse{}, utils.SystemError("failed to create year")
	}

	setting := &models.YearFormSetting{
		YearID: year.YearID,
		ClusterActive: false,
		FlowerActive : false,
		PodActive 		:		false,
		PreHarvestActive 	:	false,
		HarvestGradingActive :	false,
	}

	if err := s.yearRepo.CreateFormSetting(tx, setting); err != nil {
		tx.Rollback()
		return YearCreateResponse{}, utils.SystemError("failed to create year form setting")
	}

	if err := tx.Commit().Error; err != nil {
		return YearCreateResponse{}, utils.SystemError("failed to commit transaction")
	}
	return YearCreateResponse{
		Message: "year created successfully",
	}, nil
}