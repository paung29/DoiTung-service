package cluster

import (
	"github.com/doitung/DoiTung-service/internal/models"
	"gorm.io/gorm"
)

type ClusterRepository interface {
	CreatePole(pole *models.Pole) error
	CreateCluster(cluster *models.Cluster) error
	CreateClusterForm(form *models.ClusterForm) error
	FindPoleByZoneAndPoleNo(zoneId uint, poleNo uint) (*models.Pole, error)
	FindClusterByPoleAndClusterNo(poleId uint, clusterNo uint) (*models.Cluster, error)
	FindClusterFormByClusterId(clusterId uint) (*models.ClusterForm, error)
	UpdateCluster(db *gorm.DB, cluster *models.Cluster) error
	UpdateFormStatusByClusterId(db *gorm.DB, clusterId uint, status bool, formName string) error
	FindClusterById(clusterId uint) (*models.Cluster, error)
	GetAllClustersByPoleId(poleId uint) ([]models.Cluster, error)
	GetClusterFormByClusterId(clusterId uint) (*models.ClusterForm, error)
	GetAllClusterFormDetailsByClusterId(clusterId uint) (*models.ClusterForm, error)
	GetClusterBasicInfoByClusterId(clusterId uint) (*models.Cluster, error)
	UpdateClusterFormByClusterId(db *gorm.DB, form *models.ClusterForm) error
	GetClusterFormHistoriesByUserIdAndYearId(userId uint, yearId uint) ([]models.ClusterForm, error)
	GetAllClusterFormDetailsByZoneId(zoneId uint) ([]models.ClusterForm, error)
}
