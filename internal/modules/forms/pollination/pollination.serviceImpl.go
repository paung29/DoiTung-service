package pollination

import (
	"errors"
	"log"
	"time"

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
	yearRepo        year.YearRepository
	zoneRepo        zone.ZoneRepository
	clusterRepo     cluster.ClusterRepository
	flowerRepo      flower.FlowerRepository
	pollinationRepo PollinationRepository
}

func NewPollinationService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, flowerRepo flower.FlowerRepository, pollinationRepo PollinationRepository) PollinationService {
	return &service{
		db:              db,
		yearRepo:        yearRepo,
		zoneRepo:        zoneRepo,
		clusterRepo:     clusterRepo,
		flowerRepo:      flowerRepo,
		pollinationRepo: pollinationRepo,
	}
}

func (s *service) CreateOrUpdatePollinationForm(form PollinationFormRequest, userId uint) (PollinationFormResponse, error) {

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("year not found")
	}

	yearId := yearRecord.YearID

	// Check if the form is active in the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.PollinationActive {
		return PollinationFormResponse{}, utils.BadRequestError("pollination form is not active for this year")
	}

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneNo(yearId, int(form.ZoneNo))
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	// Check if the pole exists
	poleRecord, err := s.clusterRepo.FindPoleByZoneAndPoleNo(zoneId, form.PoleNo)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("pole not found")
	}
	poleId := poleRecord.PoleID

	// Check if the cluster exists
	clusterRecord, err := s.clusterRepo.FindClusterByPoleAndClusterNo(poleId, form.ClusterNo)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("cluster not found")
	}
	clusterId := clusterRecord.ClusterID

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
			clusterRecord.PollinationFormDone = true

			err = s.clusterRepo.UpdateCluster(tx, clusterRecord)
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
