package pod

import (
	"errors"
	"time"

	"github.com/doitung/DoiTung-service/internal/common/form"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/pollination"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db              *gorm.DB
	validator       *form.ClusterValidator
	yearRepo        year.YearRepository
	zoneRepo        zone.ZoneRepository
	clusterRepo     cluster.ClusterRepository
	pollinationRepo pollination.PollinationRepository
	podRepo         PodRepository
}

func NewPodService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, pollinationRepo pollination.PollinationRepository, podRepo PodRepository) PodService {
	validator := form.NewClusterValidator(yearRepo, zoneRepo, clusterRepo)
	return &service{
		db:              db,
		yearRepo:        yearRepo,
		zoneRepo:        zoneRepo,
		clusterRepo:     clusterRepo,
		pollinationRepo: pollinationRepo,
		podRepo:         podRepo,
		validator:       validator,
	}

}

func (s *service) CreateOrUpdatePodForm(form PodFormRequest, userId uint) (PodFormResponse, error) {

	// Validate the cluster context (year, zone, pole, cluster)
	yearId, clusterId, err := s.validator.ValidateClusterContext(
		form.Year,
		form.ZoneNo,
		form.PoleNo,
		form.ClusterNo,
		"pod",
	)
	if err != nil {
		return PodFormResponse{}, err
	}

	// Get the number of pods for the cluster

	pollinationForm, err := s.pollinationRepo.GetPollinationFormByClusterID(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormResponse{}, utils.BadRequestError("Pollination form not found for the cluster. Please submit pollination form first.")
		}
		return PodFormResponse{}, utils.SystemError("failed to get pollination form")
	}

	numberPods := pollinationForm.NumberPods
	lostPods := int(*form.LostPods)
	remainingPods := numberPods - lostPods

	// check if the pod form already exist for the cluster
	existingForm, err := s.podRepo.GetPodFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Transaction starts here
			tx := s.db.Begin()
			// Create new form
			newForm := &models.PodForm{
				ClusterID:     clusterId,
				YearID:        yearId,
				RecordedByID:  userId,
				NumberPods:    numberPods,
				LostPods:      lostPods,
				RemainingPods: remainingPods,
				RecordedDate:  time.Now(),
			}
			if err := s.podRepo.CreatePodForm(tx, newForm); err != nil {
				tx.Rollback()
				return PodFormResponse{}, err
			}

			// Update the cluster's pod form status to done
			if err := s.clusterRepo.UpdateFormStatusByClusterId(tx, clusterId, true, "pod-form"); err != nil {
				tx.Rollback()
				return PodFormResponse{}, err
			}

			if err := tx.Commit().Error; err != nil {
				return PodFormResponse{}, utils.SystemError("Failed to commit transaction")
			}

			return PodFormResponse{Message: "Pod form created successfully"}, nil
		}
		// other error
		return PodFormResponse{}, utils.SystemError("failed to check existing Pod form")
	}

	// if exist, update the form
	existingForm.NumberPods = numberPods
	existingForm.LostPods = lostPods
	existingForm.RemainingPods = remainingPods
	existingForm.RecordedByID = userId
	existingForm.RecordedDate = time.Now()

	if err := s.podRepo.UpdatePodForm(s.db, existingForm); err != nil {
		return PodFormResponse{}, utils.SystemError("failed to update Pod form")
	}

	return PodFormResponse{Message: "Pod form updated successfully"}, nil
}
