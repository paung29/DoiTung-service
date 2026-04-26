package year

import (
	"errors"
	"fmt"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

func (s *service) CreateYear(form YearCreateForm) (YearCreateResponse, error) {

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
		YearID:               year.YearID,
		ClusterActive:        false,
		FlowerActive:         false,
		PollinationActive:    false,
		PodActive:            false,
		PreHarvestActive:     false,
		HarvestGradingActive: false,
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

func (s *service) ChangeYearFormSettingStatus(form YearFormSettingStatusChange) (YearFormSettingStatusChangeResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(form.Year)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return YearFormSettingStatusChangeResponse{}, utils.NotFoundError("year not found")
		}
		return YearFormSettingStatusChangeResponse{}, utils.SystemError("failed to check year")
	}

	yearID := yearRecord.YearID

	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearID)

	if err != nil {
		return YearFormSettingStatusChangeResponse{}, utils.NotFoundError("year form setting not found")
	}

	switch form.FormName {
	case "cluster":
		yearSetting.ClusterActive = *form.ActiveStatus
	case "flower":
		yearSetting.FlowerActive = *form.ActiveStatus
	case "pollination":
		yearSetting.PollinationActive = *form.ActiveStatus
	case "pod":
		yearSetting.PodActive = *form.ActiveStatus
	case "preHarvest":
		yearSetting.PreHarvestActive = *form.ActiveStatus
	case "harvestGrading":
		yearSetting.HarvestGradingActive = *form.ActiveStatus
	default:
		return YearFormSettingStatusChangeResponse{}, utils.BadRequestError("invalid form name")
	}

	if err := s.yearRepo.UpdateFormSetting(s.db, yearSetting); err != nil {
		return YearFormSettingStatusChangeResponse{}, utils.SystemError("failed to update year form setting")
	}

	return YearFormSettingStatusChangeResponse{
		Message: "year form setting updated successfully",
	}, nil
}

func (s *service) GetYear() (GetYearResponse, error) {

	yearsModles, err := s.yearRepo.findAll()

	if err != nil {
		return  GetYearResponse{}, err
	}

	var years []string

	for _, y := range yearsModles {
		years = append(years, fmt.Sprintf("%d", y.Year))
	}

	return GetYearResponse{Years: years}, nil
}