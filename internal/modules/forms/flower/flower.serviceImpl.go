package flower

import (
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
	yearRepo    year.YearRepository
	zoneRepo    zone.ZoneRepository
	clusterRepo cluster.ClusterRepository
	flowerRepo  FlowerRepository
}

func NewFlowerService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo cluster.ClusterRepository, flowerRepo FlowerRepository) FlowerService {
	return &service{
		db:          db,
		yearRepo:    yearRepo,
		zoneRepo:    zoneRepo,
		clusterRepo: clusterRepo,
		flowerRepo:  flowerRepo,
	}
}

func (s *service) CreateFlowerForm(form FlowerFormRequest, userId uint) (FlowerFormResponse, error) {

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("year not found")
	}

	yearId := yearRecord.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.FlowerActive {
		return FlowerFormResponse{}, utils.BadRequestError("flower form is not open for this year")
	}

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneNo(uint(yearId), int(form.ZoneNo))
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	// Check if the pole exists
	poleRecord, err := s.clusterRepo.FindPoleByZoneAndPoleNo(zoneId, form.PoleNo)
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("pole not found")
	}

	// Check if the cluster exists
	clusterRecord, err := s.clusterRepo.FindClusterByPoleAndClusterNo(poleRecord.PoleID, form.ClusterNo)
	if err != nil {
		return FlowerFormResponse{}, utils.NotFoundError("cluster not found")
	}

	clusterId := clusterRecord.ClusterID

	// Create the flower form
	flowerForm := &models.FlowerForm{
		RecordedByID: userId,
		YearID:       yearId,
		ClusterID:    clusterId,
		TotalFlowers: int(form.TotalFlowers),
		Condition:    enums.Condition(form.Condition),
		Done:         true,
	}

	if err := s.flowerRepo.CreateFlowerForm(s.db, flowerForm); err != nil {
		return FlowerFormResponse{}, utils.SystemError("failed to create flower form")
	}

	return FlowerFormResponse{
		Message: "flower form created successfully",
	}, nil
}
