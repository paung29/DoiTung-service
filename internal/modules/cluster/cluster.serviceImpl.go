package cluster

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/doitung/DoiTung-service/internal/modules/year"
	"github.com/doitung/DoiTung-service/internal/modules/zone"
	"github.com/doitung/DoiTung-service/internal/types/enums"
	"github.com/doitung/DoiTung-service/internal/utils"
	"gorm.io/gorm"
)

type service struct {
	db          *gorm.DB
	accountRepo *account.AccountRepository
	yearRepo    year.YearRepository
	zoneRepo    zone.ZoneRepository
	clusterRepo ClusterRepository
}

func NewClusterService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, clusterRepo ClusterRepository) ClusterService {
	return &service{
		db:          db,
		yearRepo:    yearRepo,
		zoneRepo:    zoneRepo,
		clusterRepo: clusterRepo,
	}
}

func (s *service) CreateCluster(form ClusterCreateRequest, userId uint) (ClusterCreateResponse, error) {

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneNo(uint(form.YearId), int(form.ZoneNo))
	if err != nil {
		return ClusterCreateResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	// Transition starts here
	tx := s.db.Begin()

	// Check if the pole exists
	poleRecord, err := s.clusterRepo.FindPoleByZoneAndPoleNo(zoneId, uint(form.PoleNo))

	// If not exist, create pole
	if err != nil {
		pole := &models.Pole{
			ZoneID: zoneId,
			PoleNo: int(form.PoleNo),
		}
		if err := s.clusterRepo.CreatePole(pole); err != nil {
			tx.Rollback()
			return ClusterCreateResponse{}, utils.SystemError("Failed to create pole")
		}
		poleRecord = pole
	}
	poleId := poleRecord.PoleID
	// Check if the cluster exists
	clusterRecord, err := s.clusterRepo.FindClusterByPoleAndClusterNo(poleId, uint(form.ClusterNo))

	// If not exist, create cluster
	if err != nil {
		cluster := &models.Cluster{
			PoleID:    poleId,
			ClusterNo: int(form.ClusterNo),
		}
		if err := s.clusterRepo.CreateCluster(cluster); err != nil {
			tx.Rollback()
			return ClusterCreateResponse{}, utils.SystemError("Failed to create cluster")
		}
		clusterRecord = cluster
	}
	clusterId := clusterRecord.ClusterID

	// Check if the clusterform already exists
	existingForm, err := s.clusterRepo.FindClusterFormByClusterId(clusterId)
	if err == nil && existingForm != nil {
		tx.Rollback()
		return ClusterCreateResponse{}, utils.BadRequestError("cluster form already exists")
	}

	// Create cluster form
	clusterForm := &models.ClusterForm{
		YearID:       form.YearId,
		ClusterID:    clusterId,
		RecordedByID: userId,
		Condition:    enums.Condition(form.Condition),
	}
	if err := s.clusterRepo.CreateClusterForm(clusterForm); err != nil {
		tx.Rollback()
		return ClusterCreateResponse{}, utils.SystemError("Failed to create cluster form")
	}

	if err := tx.Commit().Error; err != nil {
		return ClusterCreateResponse{}, utils.SystemError("Failed to commit transaction")
	}

	return ClusterCreateResponse{
		Message: "cluster created successfully",
	}, nil
}
