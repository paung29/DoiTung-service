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

	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(int(form.Year))
	if err != nil {
		return ClusterCreateResponse{}, utils.NotFoundError("year not found")
	}

	yearId := yearRecord.YearID

	// Check if the form setting is open for the year
	yearSetting, err := s.yearRepo.FindFormSettingByYear(yearId)
	if err != nil {
		return ClusterCreateResponse{}, utils.NotFoundError("year setting not found")
	}

	if !yearSetting.ClusterActive {
		return ClusterCreateResponse{}, utils.BadRequestError("cluster form is not open for this year")
	}

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneNo(uint(yearId), int(form.ZoneNo))
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
		if err == gorm.ErrRecordNotFound {
			pole := &models.Pole{
				ZoneID: zoneId,
				PoleNo: int(form.PoleNo),
			}
			if err := s.clusterRepo.CreatePole(pole); err != nil {
				tx.Rollback()
				return ClusterCreateResponse{}, utils.SystemError("Failed to create pole")
			}
			poleRecord = pole
		} else {
			tx.Rollback()
			return ClusterCreateResponse{}, utils.SystemError("Failed to check pole")
		}

	}
	poleId := poleRecord.PoleID
	// Check if the cluster exists
	clusterRecord, err := s.clusterRepo.FindClusterByPoleAndClusterNo(poleId, uint(form.ClusterNo))

	// If not exist, create cluster
	if err != nil {
		// If the error is record not found
		if err == gorm.ErrRecordNotFound {
			// continue to create cluster
			cluster := &models.Cluster{
				PoleID:    poleId,
				ClusterNo: int(form.ClusterNo),
			}
			if err := s.clusterRepo.CreateCluster(cluster); err != nil {
				tx.Rollback()
				return ClusterCreateResponse{}, utils.SystemError("Failed to create cluster")
			}
			clusterRecord = cluster
		} else {
			tx.Rollback()
			return ClusterCreateResponse{}, utils.SystemError("Failed to check cluster")
		}

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
		YearID:       yearId,
		ClusterID:    clusterId,
		RecordedByID: userId,
		Condition:    enums.Condition(form.Condition),
	}
	if err := s.clusterRepo.CreateClusterForm(clusterForm); err != nil {
		tx.Rollback()
		return ClusterCreateResponse{}, utils.SystemError("Failed to create cluster form")
	}

	// Cluster form done, update cluster record
	clusterRecord.ClusterFormDone = true
	if err := s.clusterRepo.UpdateCluster(clusterRecord); err != nil {
		tx.Rollback()
		return ClusterCreateResponse{}, utils.SystemError("Failed to update cluster record")
	}

	if err := tx.Commit().Error; err != nil {
		return ClusterCreateResponse{}, utils.SystemError("Failed to commit transaction")
	}

	return ClusterCreateResponse{
		Message: "cluster created successfully",
	}, nil
}
