package harvestgrading

import (
	"errors"
	"time"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/pole"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db                 *gorm.DB
	yearRepo           year.YearRepository
	zoneRepo           zone.ZoneRepository
	poleRepo           pole.PoleRepository
	harvestGradingRepo HarvestGradingRepository
}

func NewHarvestGradingService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, poleRepo pole.PoleRepository, harvestGradingRepo HarvestGradingRepository) HarvestGradingService {
	return &service{
		db:                 db,
		yearRepo:           yearRepo,
		zoneRepo:           zoneRepo,
		poleRepo:           poleRepo,
		harvestGradingRepo: harvestGradingRepo,
	}
}

func (s *service) CreateOrUpdateHarvestGradingForm(form HarvestGradingFormRequest, userId uint) (HarvestGradingFormResponse, error) {

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		return HarvestGradingFormResponse{}, utils.NotFoundError("year not found")
	}

	yearId := yearRecord.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return HarvestGradingFormResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.HarvestGradingActive {
		return HarvestGradingFormResponse{}, utils.BadRequestError("harvest grading form is not open for this year")
	}

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneNo(uint(yearId), int(form.ZoneNo))
	if err != nil {
		return HarvestGradingFormResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	// Check if the pole exists
	poleRecord, err := s.poleRepo.GetPoleByZoneIdAndPoleNo(zoneId, form.PoleNo)
	if err != nil {
		return HarvestGradingFormResponse{}, utils.NotFoundError("pole not found")
	}
	poleId := poleRecord.PoleID

	gradeAPlusCount := int(*form.GradeAPlusCount)
	gradeAPlusWeight := int(*form.GradeAPlusWeight)
	gradeACount := int(*form.GradeACount)
	gradeAWeight := int(*form.GradeAWeight)
	gradeBCount := int(*form.GradeBCount)
	gradeBWeight := int(*form.GradeBWeight)
	gradeCCount := int(*form.GradeCCount)
	gradeCWeight := int(*form.GradeCWeight)
	gradeDPlusCount := int(*form.GradeDPlusCount)
	gradeDPlusWeight := int(*form.GradeDPlusWeight)
	undersizedCount := int(*form.UndersizedCount)
	undersizedWeight := int(*form.UndersizedWeight)
	recordedDate := time.Now()

	// Check if the harvest grading form already exists for the pole
	existingForm, err := s.harvestGradingRepo.GetHarvestGradingFormByPoleId(poleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			tx := s.db.Begin()

			// Create the harvest grading form
			harvestGradingForm := &models.HarvestGradingForm{
				YearID:           yearId,
				RecordedByID:     userId,
				PoleID:           poleId,
				GradeAPlusCount:  gradeAPlusCount,
				GradeAPlusWeight: gradeAPlusWeight,
				GradeACount:      gradeACount,
				GradeAWeight:     gradeAWeight,
				GradeBCount:      gradeBCount,
				GradeBWeight:     gradeBWeight,
				GradeCCount:      gradeCCount,
				GradeCWeight:     gradeCWeight,
				GradeDPlusCount:  gradeDPlusCount,
				GradeDPlusWeight: gradeDPlusWeight,
				UndersizedCount:  undersizedCount,
				UndersizedWeight: undersizedWeight,
				RecordedDate:     recordedDate,
			}

			// Save the harvest grading form
			if err := s.harvestGradingRepo.CreateHarvestGradingForm(tx, harvestGradingForm); err != nil {
				tx.Rollback()
				return HarvestGradingFormResponse{}, utils.SystemError("failed to create harvest grading form")
			}

			// Harvest grading form done, update pole record
			if err := s.poleRepo.UpdateHarvestGradingStatusByPoleId(poleId, true); err != nil {
				tx.Rollback()
				return HarvestGradingFormResponse{}, utils.SystemError("failed to update pole record")
			}

			// Commit the transaction
			if err := tx.Commit().Error; err != nil {
				return HarvestGradingFormResponse{}, utils.SystemError("failed to commit transaction")
			}

			return HarvestGradingFormResponse{
				Message: "harvest grading form created successfully!!!",
			}, nil
		}

		return HarvestGradingFormResponse{}, utils.NotFoundError("harvest grading form not found")
	}

	// If form does not exist, create it
	existingForm.YearID = yearId
	existingForm.RecordedByID = userId
	existingForm.GradeAPlusCount = gradeAPlusCount
	existingForm.GradeAPlusWeight = gradeAPlusWeight
	existingForm.GradeACount = gradeACount
	existingForm.GradeAWeight = gradeAWeight
	existingForm.GradeBCount = gradeBCount
	existingForm.GradeBWeight = gradeBWeight
	existingForm.GradeCCount = gradeCCount
	existingForm.GradeCWeight = gradeCWeight
	existingForm.GradeDPlusCount = gradeDPlusCount
	existingForm.GradeDPlusWeight = gradeDPlusWeight
	existingForm.UndersizedCount = undersizedCount
	existingForm.UndersizedWeight = undersizedWeight
	existingForm.RecordedDate = recordedDate

	if err := s.harvestGradingRepo.UpdateHarvestGradingForm(existingForm); err != nil {
		return HarvestGradingFormResponse{}, utils.SystemError("failed to update harvest grading form")
	}

	return HarvestGradingFormResponse{
		Message: "harvest grading form updated successfully!!!",
	}, nil
}

func (s *service) GetHarvestGradingFormDetailsByPoleID(poleId uint) (HarvestGradingFormDetails, error) {

	// Check if the pole exists
	poleRecord, err := s.poleRepo.GetPoleById(poleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return HarvestGradingFormDetails{}, utils.NotFoundError("pole not found")
		}
		return HarvestGradingFormDetails{}, utils.SystemError("failed to check pole")
	}

	harvestGradingDetails := HarvestGradingFormDetails{
		PoleId:                 poleId,
		Location:               poleRecord.Zone.ZoneName,
		PoleNo:                 uint(poleRecord.PoleNo),
		HarvestGradingFormDone: poleRecord.HarvestGradingFormDone,
	}

	harvestGradingForm, err := s.harvestGradingRepo.GetHarvestGradingFormByPoleId(poleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			harvestGradingDetails.HarvestGradingFormDone = false
			return harvestGradingDetails, nil
		}
		return HarvestGradingFormDetails{}, utils.SystemError("failed to get harvest grading form")
	}

	harvestGradingDetails.Year = harvestGradingForm.YearID
	harvestGradingDetails.GradeAPlusCount = uint(harvestGradingForm.GradeAPlusCount)
	harvestGradingDetails.GradeAPlusWeight = uint(harvestGradingForm.GradeAPlusWeight)
	harvestGradingDetails.GradeACount = uint(harvestGradingForm.GradeACount)
	harvestGradingDetails.GradeAWeight = uint(harvestGradingForm.GradeAWeight)
	harvestGradingDetails.GradeBCount = uint(harvestGradingForm.GradeBCount)
	harvestGradingDetails.GradeBWeight = uint(harvestGradingForm.GradeBWeight)
	harvestGradingDetails.GradeCCount = uint(harvestGradingForm.GradeCCount)
	harvestGradingDetails.GradeCWeight = uint(harvestGradingForm.GradeCWeight)
	harvestGradingDetails.GradeDPlusCount = uint(harvestGradingForm.GradeDPlusCount)
	harvestGradingDetails.GradeDPlusWeight = uint(harvestGradingForm.GradeDPlusWeight)
	harvestGradingDetails.UndersizedCount = uint(harvestGradingForm.UndersizedCount)
	harvestGradingDetails.UndersizedWeight = uint(harvestGradingForm.UndersizedWeight)

	return harvestGradingDetails, nil
}

func (s *service) GetHarvestGradingFormHistories(userId uint, year uint) (HarvestGradingFormHistoriesResponse, error) {

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		return HarvestGradingFormHistoriesResponse{}, utils.NotFoundError("year not found")
	}

	yearId := yearRecord.YearID

	harvestGradingForms, err := s.harvestGradingRepo.GetHarvestGradingFormsByUserIdAndYearId(s.db, userId, yearId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return HarvestGradingFormHistoriesResponse{
				HarvestGradingFormHistories: []HarvestGradingFormHistory{},
			}, nil
		}
		return HarvestGradingFormHistoriesResponse{}, utils.SystemError("failed to get harvest grading form histories")
	}

	var harvestGradingFormHistories []HarvestGradingFormHistory
	for _, form := range harvestGradingForms {
		poleRecord, err := s.poleRepo.GetPoleById(form.PoleID)
		if err != nil {
			return HarvestGradingFormHistoriesResponse{}, utils.SystemError("failed to get pole record")
		}

		harvestGradingFormHistory := HarvestGradingFormHistory{
			PoleId:                 form.PoleID,
			Location:               poleRecord.Zone.ZoneName,
			PoleNo:                 uint(poleRecord.PoleNo),
			HarvestGradingFormDone: poleRecord.HarvestGradingFormDone,
			CreatedAt:              form.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:              form.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		harvestGradingFormHistories = append(harvestGradingFormHistories, harvestGradingFormHistory)
	}

	return HarvestGradingFormHistoriesResponse{
		HarvestGradingFormHistories: harvestGradingFormHistories,
	}, nil
}
