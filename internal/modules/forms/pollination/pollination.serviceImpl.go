package pollination

import (
	"errors"
	"log"
	"time"

	"github.com/doitung/DoiTung-service/internal/common/form"
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/cluster"
	"github.com/doitung/DoiTung-service/internal/modules/forms/flower"
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
	flowerRepo      flower.FlowerRepository
	pollinationRepo PollinationRepository
}

func NewPollinationService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, flowerRepo flower.FlowerRepository, pollinationRepo PollinationRepository) PollinationService {
	validator := form.NewClusterValidator(yearRepo, zoneRepo, clusterRepo)
	return &service{
		db:              db,
		validator:       validator,
		yearRepo:        yearRepo,
		zoneRepo:        zoneRepo,
		clusterRepo:     clusterRepo,
		flowerRepo:      flowerRepo,
		pollinationRepo: pollinationRepo,
	}
}

func (s *service) CreateOrUpdatePollinationForm(form PollinationFormRequest, userId uint) (PollinationFormResponse, error) {

	yearId, clusterId, err := s.validator.ValidateClusterContext(
		form.Year,
		form.ZoneNo,
		form.PoleNo,
		form.ClusterNo,
		"pollination",
	)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("cluster not found")
	}

	// Get Flower Record
	flowerRecord, err := s.flowerRepo.GetFlowerFormByClusterID(s.db, clusterId)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("flower form must be filled before filling pollination form")
	}

	numberPods := int(*form.NumberPods)
	unsuccessfulPollination := int(*form.UnsuccessfulPollination)
	goodFlowers := numberPods + unsuccessfulPollination
	badFlowers := flowerRecord.TotalFlowers - goodFlowers
	if badFlowers < 0 {
		return PollinationFormResponse{}, utils.BadRequestError("number of pods and unsuccessful pollination cannot be greater than total flowers")
	}
	condition := enums.Condition(form.Condition)
	successMessage := ""

	// Check if the pollination form already exists for the cluster
	existingForm, err := s.pollinationRepo.GetPollinationFormByClusterID(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// Transaction Starts here
			tx := s.db.Begin()

			newForm := &models.PollinationForm{
				ClusterID:               clusterId,
				YearID:                  yearId,
				RecordedByID:            userId,
				NumberPods:              numberPods,
				UnsuccessfulPollination: unsuccessfulPollination,
				Condition:               condition,
				GoodFlowers:             goodFlowers,
				BadFlowers:              badFlowers,
				RecordedDate:            time.Now(),
			}

			err = s.pollinationRepo.CreatePollinationForm(tx, newForm)
			if err != nil {
				tx.Rollback()
				return PollinationFormResponse{}, utils.SystemError("failed to create pollination form")
			}
			successMessage = "pollination form created successfully!!"
			log.Println(successMessage)

			// Pollination form done, update cluster record
			err = s.clusterRepo.UpdateFormStatusByClusterId(tx, clusterId, true, "pollination")
			if err != nil {
				tx.Rollback()
				return PollinationFormResponse{}, utils.SystemError("failed to update cluster record")
			}

			if err := tx.Commit().Error; err != nil {
				return PollinationFormResponse{}, utils.SystemError("failed to commit transaction")
			}
			return PollinationFormResponse{Message: successMessage}, nil
		}
		return PollinationFormResponse{}, utils.SystemError("failed to check existing pollination form")
	}

	// If form exists, update it. Otherwise, create a new form.

	existingForm.ClusterID = clusterId
	existingForm.YearID = yearId
	existingForm.RecordedByID = userId
	existingForm.NumberPods = numberPods
	existingForm.UnsuccessfulPollination = unsuccessfulPollination
	existingForm.GoodFlowers = goodFlowers
	existingForm.BadFlowers = badFlowers
	existingForm.Condition = condition
	existingForm.RecordedDate = time.Now()

	// Update the pollination form
	if err = s.pollinationRepo.UpdatePollinationForm(s.db, existingForm); err != nil {
		return PollinationFormResponse{}, utils.SystemError("failed to update pollination form")
	}
	successMessage = "pollination form updated successfully"

	return PollinationFormResponse{Message: successMessage}, nil

}

func (s *service) GetPollinationFormDetails(clusterId uint) (PollinationFormDetails, error) {

	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PollinationFormDetails{}, utils.BadRequestError("cluster not found")
		}
		return PollinationFormDetails{}, utils.SystemError("failed to get cluster information")
	}

	pollinationDetails := PollinationFormDetails{
		ClusterId:           clusterInfo.ClusterID,
		Location:            clusterInfo.Pole.Zone.ZoneName,
		PoleNo:              uint(clusterInfo.Pole.PoleNo),
		ClusterNo:           uint(clusterInfo.ClusterNo),
		PollinationFormDone: clusterInfo.PollinationFormDone,
	}
	pollinationFormRecord, err := s.pollinationRepo.GetPollinationFormByClusterID(s.db, clusterId)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pollinationDetails, nil
		}
		return PollinationFormDetails{}, utils.SystemError("failed to get pollination form")

	}
	pollinationDetails.NumberPods = uint(pollinationFormRecord.NumberPods)
	pollinationDetails.UnsuccessfulPollination = uint(pollinationFormRecord.UnsuccessfulPollination)
	pollinationDetails.GoodFlowers = uint(pollinationFormRecord.GoodFlowers)
	pollinationDetails.BadFlowers = uint(pollinationFormRecord.BadFlowers)
	pollinationDetails.Condition = string(pollinationFormRecord.Condition)

	return pollinationDetails, nil
}
