package preharvest

import (
	"errors"
	"time"

	"github.com/doitung/DoiTung-service/internal/common/form"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pod"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db             *gorm.DB
	yearRepo       year.YearRepository
	zoneRepo       zone.ZoneRepository
	clusterRepo    cluster.ClusterRepository
	validator      *form.ClusterValidator
	podRepo        pod.PodRepository
	preHarvestRepo PreHarvestRepository
}

func NewPreHarvestService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, podRepo pod.PodRepository, preHarvestRepo PreHarvestRepository) PreHarvestService {
	validator := form.NewClusterValidator(yearRepo, zoneRepo, clusterRepo)
	return &service{
		db:             db,
		yearRepo:       yearRepo,
		zoneRepo:       zoneRepo,
		clusterRepo:    clusterRepo,
		validator:      validator,
		podRepo:        podRepo,
		preHarvestRepo: preHarvestRepo,
	}
}

func (s *service) CreateOrUpdatePreHarvestForm(form PreHarvestFormRequest, userId uint) (PreHarvestFormResponse, error) {

	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(form.ClusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormResponse{}, utils.BadRequestError("cluster not found")
		}
		return PreHarvestFormResponse{}, utils.SystemError("failed to get cluster information")
	}

	clusterId := clusterInfo.ClusterID
	yearId := clusterInfo.Pole.Zone.Year.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return PreHarvestFormResponse{}, utils.NotFoundError("year setting not found")
	}
	if !yearSetting.PreHarvestActive {
		return PreHarvestFormResponse{}, utils.BadRequestError("preHarvest form is not open for this year")
	}

	// Get the number of pods for the cluster
	podRecord, err := s.podRepo.GetPodFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormResponse{}, utils.BadRequestError("Pod form not found for the cluster. Please submit pod form first.")
		}
		return PreHarvestFormResponse{}, utils.SystemError("failed to get pod form")
	}

	remainPods := int(podRecord.RemainingPods)

	numberPodsSecondRound := int(*form.NumberPodsSecondRound)
	lostPodsBeforeHarvest := remainPods - numberPodsSecondRound
	if lostPodsBeforeHarvest < 0 {
		return PreHarvestFormResponse{}, utils.BadRequestError("number of pods in the second round cannot be greater than remaining pods")
	}
	removedPods := int(*form.RemovedPods)
	plantsRemoved := int(*form.PlantsRemoved)
	condition := enums.Condition(form.Condition)
	recordedDate := time.Now()

	// Check if a preHarvest form already exists for the cluster
	existingForm, err := s.preHarvestRepo.GetPreHarvestFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// transaction starts here
			tx := s.db.Begin()
			// Create a new preHarvest form
			newForm := &models.PreHarvestForm{
				ClusterID:             clusterId,
				YearID:                yearId,
				RecordedByID:          userId,
				NumberPodsSecondRound: numberPodsSecondRound,
				LostPodsBeforeHarvest: lostPodsBeforeHarvest,
				RemovedPods:           removedPods,
				PlantsRemoved:         plantsRemoved,
				Condition:             condition,
				RecordedDate:          recordedDate,
			}
			if err := tx.Create(&newForm).Error; err != nil {
				tx.Rollback()
				return PreHarvestFormResponse{}, utils.SystemError("failed to create preHarvest form")
			}

			// Update the cluster's preHarvest form status to done
			if err := s.clusterRepo.UpdateFormStatusByClusterId(tx, clusterId, true, "preHarvest"); err != nil {
				tx.Rollback()
				return PreHarvestFormResponse{}, utils.SystemError("failed to update cluster form status")
			}

			if err := tx.Commit().Error; err != nil {
				return PreHarvestFormResponse{}, utils.SystemError("failed to commit transaction")
			}

			return PreHarvestFormResponse{Message: "preHarvest form created successfully"}, nil
		} else {
			return PreHarvestFormResponse{}, utils.SystemError("failed to get existing preHarvest form")
		}
	}

	// Update the existing preHarvest form
	existingForm.NumberPodsSecondRound = numberPodsSecondRound
	existingForm.LostPodsBeforeHarvest = lostPodsBeforeHarvest
	existingForm.RemovedPods = removedPods
	existingForm.PlantsRemoved = plantsRemoved
	existingForm.Condition = condition
	existingForm.RecordedDate = recordedDate

	if err := s.preHarvestRepo.UpdatePreHarvestForm(s.db, existingForm); err != nil {
		return PreHarvestFormResponse{}, utils.SystemError("failed to update preHarvest form")
	}

	return PreHarvestFormResponse{Message: "preHarvest form updated successfully"}, nil
}

func (s *service) GetPreHarvestFormDetails(clusterId uint) (PreHarvestFormDetails, error) {

	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormDetails{}, utils.BadRequestError("cluster not found")
		}
		return PreHarvestFormDetails{}, utils.SystemError("failed to get cluster information")
	}

	podRecord, err := s.podRepo.GetPodFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormDetails{}, utils.BadRequestError("pod form not found for the cluster")
		}
		return PreHarvestFormDetails{}, utils.SystemError("failed to get pod form")
	}

	preHarvestDetails := PreHarvestFormDetails{
		ClusterId:          clusterInfo.ClusterID,
		Location:           clusterInfo.Pole.Zone.ZoneName,
		PoleNo:             uint(clusterInfo.Pole.PoleNo),
		ClusterNo:          uint(clusterInfo.ClusterNo),
		RemainingPods:      uint(podRecord.RemainingPods),
		PreHarvestFormDone: clusterInfo.PreHarvestFormDone,
	}

	preHarvestFormRecord, err := s.preHarvestRepo.GetPreHarvestFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return preHarvestDetails, nil // Return basic details with form done status
		}
		return PreHarvestFormDetails{}, utils.SystemError("failed to get preHarvest form")
	}

	preHarvestDetails.NumberPodsSecondRound = uint(preHarvestFormRecord.NumberPodsSecondRound)
	preHarvestDetails.LostPodsBeforeHarvest = uint(preHarvestFormRecord.LostPodsBeforeHarvest)
	preHarvestDetails.RemovedPods = uint(preHarvestFormRecord.RemovedPods)
	preHarvestDetails.PlantsRemoved = uint(preHarvestFormRecord.PlantsRemoved)
	preHarvestDetails.Condition = string(preHarvestFormRecord.Condition)

	return preHarvestDetails, nil
}

func (s *service) GetPreHarvestFormHistories(userId uint, year uint) (PreHarvestFormHistoriesResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormHistoriesResponse{}, utils.BadRequestError("year not found")
		}
		return PreHarvestFormHistoriesResponse{}, utils.SystemError("failed to get year information")
	}

	preHarvestFormRecords, err := s.preHarvestRepo.GetPreHarvestFormsByUserIdAndYear(s.db, userId, yearRecord.YearID)
	if err != nil {
		return PreHarvestFormHistoriesResponse{}, utils.SystemError("failed to get preHarvest form histories")
	}

	preHarvestFormHistories := make([]PreHarvestFormHistory, 0, len(preHarvestFormRecords))
	for number, record := range preHarvestFormRecords {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(record.ClusterID)
		if err != nil {
			return PreHarvestFormHistoriesResponse{}, utils.SystemError("failed to get cluster information for form history")
		}

		preHarvestFormHistories = append(preHarvestFormHistories, PreHarvestFormHistory{
			No:           uint(number + 1),
			ClusterId:    record.ClusterID,
			Location:     clusterInfo.Pole.Zone.ZoneName,
			PoleNo:       uint(clusterInfo.Pole.PoleNo),
			ClusterNo:    uint(clusterInfo.ClusterNo),
			ProgressDone: utils.CalculateClusterProgress(*clusterInfo),
			CreatedAt:    record.CreatedAt.Format("2006-01-02"),
			UpdatedAt:    record.UpdatedAt.Format("2006-01-02"),
		})
	}
	return PreHarvestFormHistoriesResponse{PreHarvestForms: preHarvestFormHistories}, nil
}

func (s *service) GetPreHarvestFormByZoneId(zoneId uint) (PreHarvestFormLists, error) {

	zoneRecord, err := s.zoneRepo.FindById(zoneId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PreHarvestFormLists{}, utils.BadRequestError("zone not found")
		}
		return PreHarvestFormLists{}, utils.SystemError("failed to get preHarvest forms by zone id")
	}

	zoneId = zoneRecord.ZoneID

	preHarvestForms, err := s.preHarvestRepo.GetPreHarvestFormsByZoneId(s.db, zoneId)
	if err != nil {
		return PreHarvestFormLists{}, utils.SystemError("failed to get preHarvest forms by zone id")
	}

	preHarvestFormDetailsList := make([]PreHarvestFormDetails, 0, len(preHarvestForms))
	for number, record := range preHarvestForms {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(record.ClusterID)
		if err != nil {
			return PreHarvestFormLists{}, utils.SystemError("failed to get cluster information for form list")
		}

		remainPods := record.LostPodsBeforeHarvest + record.NumberPodsSecondRound
		preHarvestFormDetailsList = append(preHarvestFormDetailsList, PreHarvestFormDetails{
			No:                    uint(number + 1),
			ClusterId:             record.ClusterID,
			Location:              clusterInfo.Pole.Zone.ZoneName,
			PoleNo:                uint(clusterInfo.Pole.PoleNo),
			ClusterNo:             uint(clusterInfo.ClusterNo),
			RemainingPods:         uint(remainPods),
			NumberPodsSecondRound: uint(record.NumberPodsSecondRound),
			LostPodsBeforeHarvest: uint(record.LostPodsBeforeHarvest),
			RemovedPods:           uint(record.RemovedPods),
			PlantsRemoved:         uint(record.PlantsRemoved),
			Condition:             string(record.Condition),
			PreHarvestFormDone:    clusterInfo.PreHarvestFormDone,
			RecordedBy:            record.RecordedBy.Name,
			Date:                  record.UpdatedAt.Format("2006-01-02"),
		})
	}

	return PreHarvestFormLists{PreHarvestForms: preHarvestFormDetailsList}, nil
}
