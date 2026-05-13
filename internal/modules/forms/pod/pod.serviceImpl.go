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
	"github.com/doitung/DoiTung-service/internal/types/enums"
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

	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(form.ClusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormResponse{}, utils.BadRequestError("cluster not found")
		}
		return PodFormResponse{}, utils.SystemError("failed to get cluster information")
	}

	clusterId := clusterInfo.ClusterID
	yearId := clusterInfo.Pole.Zone.Year.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return PodFormResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.PodActive {
		return PodFormResponse{}, utils.BadRequestError("pod form is not open")
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
	if remainingPods < 0 {
		return PodFormResponse{}, utils.BadRequestError("number of lost pods cannot be greater than number of pods in pollination form")
	}
	condition := enums.Condition(form.Condition)

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
				Condition:     condition,
			}
			if err := s.podRepo.CreatePodForm(tx, newForm); err != nil {
				tx.Rollback()
				return PodFormResponse{}, err
			}

			// Update the cluster's pod form status to done
			if err := s.clusterRepo.UpdateFormStatusByClusterId(tx, clusterId, true, "pod"); err != nil {
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
	existingForm.Condition = condition
	existingForm.RecordedDate = time.Now()

	if err := s.podRepo.UpdatePodForm(s.db, existingForm); err != nil {
		return PodFormResponse{}, utils.SystemError("failed to update Pod form")
	}

	return PodFormResponse{Message: "Pod form updated successfully"}, nil
}

func (s *service) GetPodFormDetails(clusterId uint) (PodFormDetails, error) {
	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormDetails{}, utils.BadRequestError("cluster not found")
		}
		return PodFormDetails{}, utils.SystemError("failed to get cluster information")
	}

	clusterId = clusterInfo.ClusterID

	// Get the pollination data
	pollinationRecord, err := s.pollinationRepo.GetPollinationFormByClusterID(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormDetails{}, utils.BadRequestError("pollination form not found for the cluster. Please submit pollination form first.")
		}
		return PodFormDetails{}, utils.SystemError("failed to get pollination form")
	}

	podDetails := PodFormDetails{
		ClusterId:   clusterInfo.ClusterID,
		Location:    clusterInfo.Pole.Zone.ZoneName,
		PoleNo:      uint(clusterInfo.Pole.PoleNo),
		ClusterNo:   uint(clusterInfo.ClusterNo),
		NumberPods:  uint(pollinationRecord.NumberPods),
		PodFormDone: clusterInfo.PodFormDone,
	}

	podFormRecord, err := s.podRepo.GetPodFormByClusterId(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return podDetails, nil // return basic details with PodFormDone = false
		}
		return PodFormDetails{}, utils.SystemError("failed to get pod form details")
	}

	podDetails.LostPods = uint(podFormRecord.LostPods)
	podDetails.RemainingPods = uint(podFormRecord.RemainingPods)
	podDetails.Condition = string(podFormRecord.Condition)

	return podDetails, nil
}

func (s *service) GetPodFormHistories(userId uint, year uint) (PodFormHistoriesResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormHistoriesResponse{}, utils.BadRequestError("year not found")
		}
		return PodFormHistoriesResponse{}, utils.SystemError("failed to get year information")
	}

	podFormHistories, err := s.podRepo.GetPodFormHistoriesByUserIdAndYearId(s.db, userId, yearRecord.YearID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PodFormHistoriesResponse{PodFormHistories: []cluster.ClusterInfo{}}, nil
		}
		return PodFormHistoriesResponse{}, utils.SystemError("failed to get pod form histories")
	}

	var podFormHistoriesResponse []cluster.ClusterInfo
	for number, history := range podFormHistories {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(history.ClusterID)
		if err != nil {
			return PodFormHistoriesResponse{}, utils.SystemError("failed to get cluster information for pod form history")
		}
		clusterProgress := utils.CalculateClusterProgress(*clusterInfo)
		podFormHistoriesResponse = append(podFormHistoriesResponse, cluster.ClusterInfo{
			No:           number + 1,
			ClusterId:    history.ClusterID,
			Location:     clusterInfo.Pole.Zone.ZoneName,
			PoleNo:       clusterInfo.Pole.PoleNo,
			ClusterNo:    clusterInfo.ClusterNo,
			ProgressDone: int(clusterProgress),
			CreatedAt:    history.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:    history.UpdatedAt.Format("2006-01-02 15:04"),
		})
	}

	return PodFormHistoriesResponse{PodFormHistories: podFormHistoriesResponse}, nil
}
