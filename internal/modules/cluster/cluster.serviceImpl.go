package cluster

import (
	"time"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/doitung/DoiTung-service/internal/modules/account"
	"github.com/doitung/DoiTung-service/internal/modules/pole"
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
	poleRepo    pole.PoleRepository
	clusterRepo ClusterRepository
}

func NewClusterService(db *gorm.DB, yearRepo year.YearRepository, zoneRepo zone.ZoneRepository, poleRepo pole.PoleRepository, clusterRepo ClusterRepository) ClusterService {
	return &service{
		db:          db,
		yearRepo:    yearRepo,
		zoneRepo:    zoneRepo,
		poleRepo:    poleRepo,
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
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneId(uint(yearId), int(form.ZoneId))
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
		RecordedDate: time.Now(),
	}
	if err := s.clusterRepo.CreateClusterForm(clusterForm); err != nil {
		tx.Rollback()
		return ClusterCreateResponse{}, utils.SystemError("Failed to create cluster form")
	}

	// Cluster form done, update cluster record
	clusterRecord.ClusterFormDone = true
	if err := s.clusterRepo.UpdateCluster(tx, clusterRecord); err != nil {
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

func (s *service) GetClustersByZone(year int, zoneNo int) (ClustersByZoneResponse, error) {
	// Check if the year exists
	yearRecord, err := s.yearRepo.FindByYear(year)
	if err != nil {
		return ClustersByZoneResponse{}, utils.NotFoundError("year not found")
	}
	yearId := yearRecord.YearID

	// Check if the zone exists
	zoneRecord, err := s.zoneRepo.FindByYearAndZoneId(yearId, zoneNo)
	if err != nil {
		return ClustersByZoneResponse{}, utils.NotFoundError("zone not found")
	}
	zoneId := zoneRecord.ZoneID

	// Get all poles by zone id
	poles, err := s.poleRepo.GetAllPolesByZoneId(zoneId)
	if err != nil {
		return ClustersByZoneResponse{}, utils.SystemError("Failed to get poles by zone id")
	}

	var clusters []models.Cluster
	for _, pole := range poles {
		poleClusters, err := s.clusterRepo.GetAllClustersByPoleId(pole.PoleID)
		if err != nil {
			return ClustersByZoneResponse{}, utils.SystemError("Failed to get clusters by pole id")
		}
		clusters = append(clusters, poleClusters...)
	}
	var clusterResponses []ClusterInfo

	for i, cluster := range clusters {
		clusterProgress := utils.CalculateClusterProgress(cluster)

		clusterResponses = append(clusterResponses, ClusterInfo{
			No:           i + 1,
			ClusterId:    cluster.ClusterID,
			PoleNo:       cluster.Pole.PoleNo,
			ClusterNo:    cluster.ClusterNo,
			Location:     cluster.Pole.Zone.ZoneName,
			CreatedAt:    cluster.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:    cluster.UpdatedAt.Format("2006-01-02 15:04"),
			ProgressDone: int(clusterProgress),
		})

	}

	return ClustersByZoneResponse{Clusters: clusterResponses}, nil
}

func (s *service) GetClusterFormByClusterId(clusterId int) (ClusterFormResponse, error) {

	clusterRecord, err := s.clusterRepo.GetAllClusterFormDetailsByClusterId(uint(clusterId))
	if err != nil {
		return ClusterFormResponse{}, utils.NotFoundError("cluster form not found")
	}

	var clusterFormResponse ClusterFormResponse
	clusterFormResponse.ClusterId = clusterRecord.ClusterID
	clusterFormResponse.Location = clusterRecord.Cluster.Pole.Zone.ZoneName
	clusterFormResponse.ClusterId = clusterRecord.ClusterID
	clusterFormResponse.PoleNo = clusterRecord.Cluster.Pole.PoleNo
	clusterFormResponse.ClusterNo = clusterRecord.Cluster.ClusterNo
	clusterFormResponse.Condition = string(clusterRecord.Condition)

	return clusterFormResponse, nil
}

func (s *service) UpdateClusterForm(form ClusterUpdateRequest) (ClusterUpdateResponse, error) {

	clusterId := form.ClusterId
	// Check if the cluster form exists
	existingForm, err := s.clusterRepo.FindClusterFormByClusterId(clusterId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ClusterUpdateResponse{}, utils.NotFoundError("cluster form not found")
		}
		return ClusterUpdateResponse{}, utils.SystemError("failed to check cluster form")
	}

	existingForm.Condition = enums.Condition(form.Condition)

	if err := s.clusterRepo.UpdateClusterFormByClusterId(s.db, existingForm); err != nil {
		return ClusterUpdateResponse{}, utils.SystemError("failed to update cluster form")
	}

	return ClusterUpdateResponse{
		Message: "cluster form updated successfully!!!",
	}, nil
}

func (s *service) GetClusterFormHistories(userId uint, year uint) (ClusterFormHistoriesResponse, error) {

	yearModel, err := s.yearRepo.FindByYear(int(year))
	if err != nil {
		return ClusterFormHistoriesResponse{}, utils.NotFoundError("year not found")
	}

	clusterFormHistories, err := s.clusterRepo.GetClusterFormHistoriesByUserIdAndYearId(userId, yearModel.YearID)
	if err != nil {
		return ClusterFormHistoriesResponse{}, utils.SystemError("failed to get cluster form histories")
	}

	var clusterFormHistoryResponses []ClusterInfo
	for number, history := range clusterFormHistories {
		progressDone := utils.CalculateClusterProgress(history.Cluster)
		clusterFormHistoryResponses = append(clusterFormHistoryResponses, ClusterInfo{
			No:           number + 1,
			ClusterId:    history.ClusterID,
			Location:     history.Cluster.Pole.Zone.ZoneName,
			PoleNo:       history.Cluster.Pole.PoleNo,
			ProgressDone: int(progressDone),
			ClusterNo:    history.Cluster.ClusterNo,
			CreatedAt:    history.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:    history.UpdatedAt.Format("2006-01-02 15:04"),
		})
	}

	return ClusterFormHistoriesResponse{
		ClusterFormHistories: clusterFormHistoryResponses,
	}, nil
}

func (s *service) GetAllClustersFormByZone(zoneId uint) (GetAllClustersFormByZoneResponse, error) {

	zoneRecord, err := s.zoneRepo.FindById(zoneId)
	if err != nil {
		return GetAllClustersFormByZoneResponse{}, utils.NotFoundError("zone not found")
	}

	zoneId = zoneRecord.ZoneID

	clusterForms, err := s.clusterRepo.GetAllClusterFormDetailsByZoneId(zoneId)
	if err != nil {
		return GetAllClustersFormByZoneResponse{}, utils.SystemError("failed to get cluster forms by zone id")
	}

	var clusterFormDetails []ClusterFormDetails
	for i, clusterForm := range clusterForms {
		clusterFormDetails = append(clusterFormDetails, ClusterFormDetails{
			No:         i + 1,
			Date:       clusterForm.CreatedAt.Format("2006-01-02"),
			ClusterId:  clusterForm.ClusterID,
			PoleNo:     clusterForm.Cluster.Pole.PoleNo,
			ClusterNo:  clusterForm.Cluster.ClusterNo,
			Condition:  string(clusterForm.Condition),
			RecordedBy: clusterForm.RecordedBy.Name,
		})
	}

	return GetAllClustersFormByZoneResponse{
		ClusterForms: clusterFormDetails,
	}, nil
}
