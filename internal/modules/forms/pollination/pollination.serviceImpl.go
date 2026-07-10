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

	clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(form.ClusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PollinationFormResponse{}, utils.BadRequestError("cluster not found")
		}
		return PollinationFormResponse{}, utils.SystemError("failed to get cluster information")
	}

	clusterId := clusterInfo.ClusterID
	yearId := clusterInfo.Pole.Zone.Year.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return PollinationFormResponse{}, utils.NotFoundError("year setting not found")
	}
	if !yearSetting.PollinationActive {
		return PollinationFormResponse{}, utils.BadRequestError("pollination form is not open for this year")
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

	flowerRecord, err := s.flowerRepo.GetFlowerFormByClusterID(s.db, clusterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PollinationFormDetails{}, utils.BadRequestError("pollination form not found for the cluster")
		}
		return PollinationFormDetails{}, utils.SystemError("failed to get pollination form")
	}

	pollinationDetails := PollinationFormDetails{
		ClusterId:           clusterInfo.ClusterID,
		Location:            clusterInfo.Pole.Zone.ZoneName,
		PoleNo:              uint(clusterInfo.Pole.PoleNo),
		ClusterNo:           uint(clusterInfo.ClusterNo),
		TotalFlowers:        uint(flowerRecord.TotalFlowers),
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

func (s *service) GetPollinationFormHistories(userId uint, year uint) (PollinationFormHistoriesResponse, error) {

	yearRecord, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PollinationFormHistoriesResponse{}, utils.BadRequestError("year not found")
		}
		return PollinationFormHistoriesResponse{}, utils.SystemError("failed to get year information")
	}

	pollinationFormHistories, err := s.pollinationRepo.GetPollinationFormHistoriesByUserIdAndYearId(s.db, userId, yearRecord.YearID)
	if err != nil {
		return PollinationFormHistoriesResponse{}, utils.SystemError("failed to get pollination form histories")
	}

	pollinationFormHistoriesResponse := make([]PollinationFormHistory, 0, len(pollinationFormHistories))
	for _, history := range pollinationFormHistories {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(history.ClusterID)
		if err != nil {
			return PollinationFormHistoriesResponse{}, utils.SystemError("failed to get cluster information for pollination form history")
		}

		clusterProgress := utils.CalculateClusterProgress(*clusterInfo)
		pollinationFormHistoriesResponse = append(pollinationFormHistoriesResponse, PollinationFormHistory{
			ClusterId:    history.ClusterID,
			Location:     clusterInfo.Pole.Zone.ZoneName,
			PoleNo:       uint(clusterInfo.Pole.PoleNo),
			ClusterNo:    uint(clusterInfo.ClusterNo),
			ProgressDone: clusterProgress,
			CreatedAt:    history.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:    history.UpdatedAt.Format("2006-01-02 15:04"),
		})
	}
	return PollinationFormHistoriesResponse{PollinationFormHistories: pollinationFormHistoriesResponse}, nil
}

func (s *service) GetPollinationFormsByZoneId(zoneId uint) (PollinationFormLists, error) {

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindById(zoneId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return PollinationFormLists{}, utils.BadRequestError("zone not found")
		}
		return PollinationFormLists{}, utils.SystemError("failed to get zone information")
	}

	zoneId = zoneRecord.ZoneID

	pollinationForms, err := s.pollinationRepo.GetPollinationFormsByZoneId(s.db, zoneId)
	if err != nil {
		return PollinationFormLists{}, utils.SystemError("failed to get pollination forms by zone id")
	}

	pollinationFormDetailsList := make([]PollinationFormDetails, 0, len(pollinationForms))
	for i, form := range pollinationForms {
		clusterInfo, err := s.clusterRepo.GetClusterBasicInfoByClusterId(form.ClusterID)
		if err != nil {
			return PollinationFormLists{}, utils.SystemError("failed to get cluster information for pollination form list")
		}

		totalFlowers := form.GoodFlowers + form.BadFlowers
		pollinationFormDetailsList = append(pollinationFormDetailsList, PollinationFormDetails{
			No:                      i + 1,
			ClusterId:               form.ClusterID,
			Location:                clusterInfo.Pole.Zone.ZoneName,
			PoleNo:                  uint(clusterInfo.Pole.PoleNo),
			ClusterNo:               uint(clusterInfo.ClusterNo),
			TotalFlowers:            uint(totalFlowers),
			NumberPods:              uint(form.NumberPods),
			UnsuccessfulPollination: uint(form.UnsuccessfulPollination),
			GoodFlowers:             uint(form.GoodFlowers),
			BadFlowers:              uint(form.BadFlowers),
			Condition:               string(form.Condition),
			PollinationFormDone:     clusterInfo.PollinationFormDone,
			Date:                    form.UpdatedAt.Format("2006-01-02"),
			RecordedBy:              form.RecordedBy.Name,
		})

	}

	return PollinationFormLists{PollinationForms: pollinationFormDetailsList}, nil
}
