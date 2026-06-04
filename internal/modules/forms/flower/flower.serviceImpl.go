package flower

import (
	"errors"
	"time"

	"github.com/doitung/DoiTung-service/internal/common/form"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db          *gorm.DB
	validator   *form.ClusterValidator
	yearRepo    year.YearRepository
	zoneRepo    zone.ZoneRepository
	clusterRepo cluster.ClusterRepository
	flowerRepo  FlowerRepository
}

func NewFlowerService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, flowerRepo FlowerRepository) FlowerService {
	validator := form.NewClusterValidator(yearRepo, zoneRepo, clusterRepo)
	return &service{
		db:          db,
		validator:   validator,
		yearRepo:    yearRepo,
		zoneRepo:    zoneRepo,
		clusterRepo: clusterRepo,
		flowerRepo:  flowerRepo,
	}
}

func (s *service) CreateOrUpdateFlowerForm(form FlowerFormRequest, userId uint) (FlowerFormResponse, error) {

	cluserInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(form.ClusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return FlowerFormResponse{}, utils.BadRequestError("cluster not found")
		}
		return FlowerFormResponse{}, utils.SystemError("failed to get cluster information")
	}

	clusterId := cluserInfo.ClusterID
	yearId := cluserInfo.Pole.Zone.Year.YearID
	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.FlowerActive {
		return FlowerFormResponse{}, utils.BadRequestError("flower form is not open")
	}

	// Check if the flower form already exists for the cluster
	existingForm, err := s.flowerRepo.GetFlowerFormByClusterID(s.db, clusterId)

	// If form does not exist, create it
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// transaction starts here
			tx := s.db.Begin()

			// Create the flower form
			flowerForm := &models.FlowerForm{
				RecordedByID: userId,
				YearID:       yearId,
				ClusterID:    clusterId,
				TotalFlowers: int(*form.TotalFlowers),
				Condition:    enums.Condition(form.Condition),
				Done:         true,
				RecordedDate: time.Now(),
			}
			// Save the flower form
			if err := s.flowerRepo.CreateFlowerForm(s.db, flowerForm); err != nil {
				tx.Rollback()
				return FlowerFormResponse{}, utils.SystemError("failed to create flower form")
			}

			// Flower form done, update cluster record
			if err := s.clusterRepo.UpdateFormStatusByClusterId(tx, clusterId, true, "flower"); err != nil {
				tx.Rollback()
				return FlowerFormResponse{}, utils.SystemError("failed to update cluster record")
			}

			// Commit the transaction
			if err := tx.Commit().Error; err != nil {
				return FlowerFormResponse{}, utils.SystemError("failed to commit transaction")
			}

			return FlowerFormResponse{
				Message: "flower form created successfully!!!",
			}, nil
		}
		// Other errors (database errors, etc.)
		return FlowerFormResponse{}, utils.SystemError("failed to check existing flower form")
	}

	// If the form already exists, update it
	existingForm.TotalFlowers = int(*form.TotalFlowers)
	existingForm.Condition = enums.Condition(form.Condition)
	existingForm.RecordedByID = userId
	existingForm.Done = true
	existingForm.RecordedDate = time.Now()

	if err := s.flowerRepo.UpdateFlowerForm(s.db, existingForm); err != nil {
		return FlowerFormResponse{}, utils.SystemError("failed to update flower form")
	}

	return FlowerFormResponse{
		Message: "flower form updated successfully",
	}, nil
}

func (s *service) GetFlowerFormDetailsByClusterID(clusterId uint) (FlowerFormDetails, error) {
	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return FlowerFormDetails{}, utils.BadRequestError("cluster not found")
		}
		return FlowerFormDetails{}, utils.SystemError("failed to get cluster information")
	}

	flowerDetails := FlowerFormDetails{
		ClusterId:      clusterInfo.ClusterID,
		Location:       clusterInfo.Pole.Zone.ZoneName,
		PoleNo:         clusterInfo.Pole.PoleNo,
		ClusterNo:      clusterInfo.ClusterNo,
		FlowerFormDone: clusterInfo.FlowerFormDone,
	}

	flowerFormRecord, err := s.flowerRepo.GetFlowerFormByClusterID(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return flowerDetails, nil
		}
		return FlowerFormDetails{}, utils.SystemError("failed to get flower form record")
	}
	flowerDetails.TotalFlowers = uint(flowerFormRecord.TotalFlowers)
	flowerDetails.Condition = string(flowerFormRecord.Condition)

	return flowerDetails, nil
}

func (s *service) GetFlowerFormHistories(userId uint, year uint) (FlowerFormHistoriesResponse, error) {
	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return FlowerFormHistoriesResponse{}, utils.BadRequestError("year not found")
		}
		return FlowerFormHistoriesResponse{}, utils.SystemError("failed to get year information")
	}
	flowerFormRecords, err := s.flowerRepo.GetFlowerFormHistoriesByUserIdAndYearId(s.db, userId, yearRecord.YearID)
	if err != nil {
		return FlowerFormHistoriesResponse{}, utils.SystemError("failed to get flower form histories")
	}

	var flowerFormHistories []FlowerFormHistory
	for _, record := range flowerFormRecords {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(record.ClusterID)
		if err != nil {
			return FlowerFormHistoriesResponse{}, utils.SystemError("failed to get cluster information")
		}
		clusterProgress := utils.CalculateClusterProgress(*clusterInfo)
		history := FlowerFormHistory{
			ClusterId:    record.ClusterID,
			Location:     clusterInfo.Pole.Zone.ZoneName,
			PoleNo:       clusterInfo.Pole.PoleNo,
			ClusterNo:    clusterInfo.ClusterNo,
			ProgressDone: clusterProgress,
			CreatedAt:    record.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    record.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		flowerFormHistories = append(flowerFormHistories, history)
	}

	return FlowerFormHistoriesResponse{
		FlowerFormHistories: flowerFormHistories,
	}, nil
}

func (s *service) GetFlowerFormsByZoneId(zoneId uint) (FlowerFormLists, error) {

	zoneRecord, err := s.zoneRepo.FindById(zoneId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return FlowerFormLists{}, utils.BadRequestError("zone not found")
		}
		return FlowerFormLists{}, utils.SystemError("failed to get zone information")
	}

	zoneId = zoneRecord.ZoneID

	flowerForms, err := s.flowerRepo.GetFlowerFormsByZoneId(s.db, zoneId)
	if err != nil {
		return FlowerFormLists{}, utils.SystemError("failed to get flower forms by zone id")
	}

	var flowerFormDetailsList []FlowerFormDetails
	for i, form := range flowerForms {
		flowerFormDetails := FlowerFormDetails{
			No:           i + 1,
			ClusterId:    form.ClusterID,
			Location:     form.Cluster.Pole.Zone.ZoneName,
			PoleNo:       form.Cluster.Pole.PoleNo,
			ClusterNo:    form.Cluster.ClusterNo,
			TotalFlowers: uint(form.TotalFlowers),
			Condition:    string(form.Condition),
			RecordedBy:   form.RecordedBy.Name,
			Date:         form.UpdatedAt.Format("2006-01-02"),
		}
		flowerFormDetailsList = append(flowerFormDetailsList, flowerFormDetails)
	}

	return FlowerFormLists{
		FlowerForms: flowerFormDetailsList,
	}, nil
}
