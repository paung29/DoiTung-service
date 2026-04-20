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
		clusterRepo:    clusterRepo,
		validator:      validator,
		podRepo:        podRepo,
		preHarvestRepo: preHarvestRepo,
	}
}

func (s *service) CreateOreUpdatePreHarvestForm(form PreHarvestFormRequest, userId uint) (PreHarvestFormResponse, error) {

	// Validate the cluster context (year, zone, pole, cluster)
	yearId, clusterId, err := s.validator.ValidateClusterContext(
		form.Year,
		form.ZoneId,
		form.PoleId,
		form.ClusterId,
		"preHarvest",
	)
	if err != nil {
		return PreHarvestFormResponse{}, utils.SystemError("Cannot validate the cluster")
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
